package cmd

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"

	pb "github.com/kamilsk/click/pkg/server/grpc"

	"github.com/kamilsk/click/pkg/config"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

type schema struct {
	Kind    string                 `yaml:"kind"`
	Payload map[string]interface{} `yaml:"payload"`
}

type factory map[*cobra.Command]map[string]func() interface{}

func (f factory) new(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("filename").Value.String())
	if err != nil {
		return nil, err
	}
	builder, ok := f[cmd][data.Kind]
	if !ok {
		return nil, errors.Errorf("unknown payload type %q", data.Kind)
	}
	entity := builder()
	if err = mapstructure.Decode(data.Payload, &entity); err != nil {
		return nil, errors.Wrapf(err, "trying to decode payload to %#v", entity)
	}
	return entity, nil
}

func (f factory) data(name string) (schema, error) {
	var (
		err error
		out schema
		raw []byte
		src io.Reader = os.Stdin
	)
	if name != "" {
		if src, err = os.Open(name); err != nil {
			return out, errors.Wrapf(err, "trying to open file %q", name)
		}
	} else {
		name = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return out, errors.Wrapf(err, "trying to read file %q", name)
	}
	err = yaml.Unmarshal(raw, &out)
	return out, errors.Wrapf(err, "trying to decode file %q as YAML", name)
}

func call(cnf config.GRPCConfig, entity interface{}) (interface{}, error) {
	conn, err := grpc.Dial(cnf.Interface, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "trying to connect to the gRPC server %q", cnf.Interface)
	}
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	switch request := entity.(type) {
	case *pb.CreateNamespaceRequest:
		client := pb.NewNamespaceClient(conn)
		return client.Create(ctx, request)
	case *pb.CreateLinkRequest:
		client := pb.NewLinkClient(conn)
		return client.Create(ctx, request)
	case *pb.CreateAliasRequest:
		client := pb.NewAliasClient(conn)
		return client.Create(ctx, request)
	case *pb.CreateTargetRequest:
		client := pb.NewTargetClient(conn)
		return client.Create(ctx, request)
	case *pb.ReadNamespaceRequest:
		client := pb.NewNamespaceClient(conn)
		return client.Read(ctx, request)
	case *pb.ReadLinkRequest:
		client := pb.NewLinkClient(conn)
		return client.Read(ctx, request)
	case *pb.ReadAliasRequest:
		client := pb.NewAliasClient(conn)
		return client.Read(ctx, request)
	case *pb.ReadTargetRequest:
		client := pb.NewTargetClient(conn)
		return client.Read(ctx, request)
	case *pb.UpdateNamespaceRequest:
		client := pb.NewNamespaceClient(conn)
		return client.Update(ctx, request)
	case *pb.UpdateLinkRequest:
		client := pb.NewLinkClient(conn)
		return client.Update(ctx, request)
	case *pb.UpdateAliasRequest:
		client := pb.NewAliasClient(conn)
		return client.Update(ctx, request)
	case *pb.UpdateTargetRequest:
		client := pb.NewTargetClient(conn)
		return client.Update(ctx, request)
	case *pb.DeleteNamespaceRequest:
		client := pb.NewNamespaceClient(conn)
		return client.Delete(ctx, request)
	case *pb.DeleteLinkRequest:
		client := pb.NewLinkClient(conn)
		return client.Delete(ctx, request)
	case *pb.DeleteAliasRequest:
		client := pb.NewAliasClient(conn)
		return client.Delete(ctx, request)
	case *pb.DeleteTargetRequest:
		client := pb.NewTargetClient(conn)
		return client.Delete(ctx, request)
	default:
		return nil, errors.Errorf("unknown type %T", request)
	}
}

func convert(entity interface{}) map[string]interface{} {
	t := reflect.ValueOf(entity).Elem().Type()
	m := make(map[string]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v, ok := f.Tag.Lookup("json")
		if !ok {
			continue
		}
		p := strings.Split(v, ",")
		if p[0] == "-" {
			continue
		}
		switch f.Type.String() {
		case "[]uint8":
			m[p[0]] = "binary"
		default:
			m[p[0]] = f.Type.String()
		}
	}
	return m
}

func edit(cmd *cobra.Command, _ []string) error {
	entity, err := entities.new(cmd)
	if err != nil {
		return err
	}
	response, err := call(cnf.Union.GRPCConfig, entity)
	if err != nil {
		return err
	}
	cmd.Printf("`ctl %s` was called, in:%#v out:%#v\n", cmd.Use, entity, response)
	return nil
}

var entities factory

var (
	controlCmd = &cobra.Command{Use: "ctl", Short: "Communicate with Click! server via gRPC"}
	createCmd  = &cobra.Command{Use: "create", Short: "Create some kind", RunE: edit}
	readCmd    = &cobra.Command{Use: "read", Short: "Read some kind", RunE: edit}
	updateCmd  = &cobra.Command{Use: "update", Short: "Update some kind", RunE: edit}
	deleteCmd  = &cobra.Command{Use: "delete", Short: "Delete some kind", RunE: edit}
	schemaCmd  = &cobra.Command{
		Use:   "schema",
		Short: "Print schema of another control command",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				target   *cobra.Command
				builders map[string]func() interface{}
				found    bool
			)
			use := cmd.Flag("for").Value.String()
			for target, builders = range entities {
				if strings.EqualFold(target.Use, use) {
					found = true
					break
				}
			}
			if !found {
				return errors.Errorf("unknown control command %q", use)
			}
			for kind, builder := range builders {
				yaml.NewEncoder(cmd.OutOrStdout()).Encode(schema{Kind: kind, Payload: convert(builder())})
				cmd.Println()
			}
			return nil
		},
	}
)

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("click_token") },
		func() error { return v.BindEnv("grpc_host") },
		func() error {
			v.SetDefault("click_token", defaults["click_token"])
			v.SetDefault("grpc_host", defaults["grpc_host"])
			return nil
		},
		func() error {
			name := ""
			flags := controlCmd.PersistentFlags()
			flags.StringVarP(&name, "filename", "f", name, "entity source (default is standard input)")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")
			flags.StringVarP((*string)(&cnf.Union.GRPCConfig.Token),
				"token", "", v.GetString("click_token"), "user access token")
			schemaCmd.Flags().String("for", "", "which command: create, read, update or delete")
			return schemaCmd.MarkFlagRequired("for")
		},
		func() error {
			entities = factory{
				createCmd: {
					"Namespace": func() interface{} { return &pb.CreateNamespaceRequest{} },
					"Link":      func() interface{} { return &pb.CreateLinkRequest{} },
					"Alias":     func() interface{} { return &pb.CreateAliasRequest{} },
					"Target":    func() interface{} { return &pb.CreateTargetRequest{} },
				},
				readCmd: {
					"Namespace": func() interface{} { return &pb.ReadNamespaceRequest{} },
					"Link":      func() interface{} { return &pb.ReadLinkRequest{} },
					"Alias":     func() interface{} { return &pb.ReadAliasRequest{} },
					"Target":    func() interface{} { return &pb.ReadTargetRequest{} },
				},
				updateCmd: {
					"Namespace": func() interface{} { return &pb.UpdateNamespaceRequest{} },
					"Link":      func() interface{} { return &pb.UpdateLinkRequest{} },
					"Alias":     func() interface{} { return &pb.UpdateAliasRequest{} },
					"Target":    func() interface{} { return &pb.UpdateTargetRequest{} },
				},
				deleteCmd: {
					"Namespace": func() interface{} { return &pb.DeleteNamespaceRequest{} },
					"Link":      func() interface{} { return &pb.DeleteLinkRequest{} },
					"Alias":     func() interface{} { return &pb.DeleteAliasRequest{} },
					"Target":    func() interface{} { return &pb.DeleteTargetRequest{} },
				},
			}
			return nil
		},
	)
	controlCmd.AddCommand(createCmd, readCmd, updateCmd, deleteCmd, schemaCmd)
}
