package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/kamilsk/click/pkg/config"
	pb "github.com/kamilsk/click/pkg/server/grpc"
	"github.com/kamilsk/click/pkg/server/grpc/middleware"
	kit "github.com/kamilsk/go-kit/pkg/strings"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	yaml "gopkg.in/yaml.v2"
)

const (
	linkKind      kind = "Link"
	namespaceKind kind = "Namespace"
	aliasKind     kind = "Alias"
	targetKind    kind = "Target"
)

var entities factory

func init() {
	entities = factory{
		createCmd: {
			linkKind:      func() interface{} { return &pb.CreateLinkRequest{} },
			namespaceKind: func() interface{} { return &pb.CreateNamespaceRequest{} },
			aliasKind:     func() interface{} { return &pb.CreateAliasRequest{} },
			targetKind:    func() interface{} { return &pb.CreateTargetRequest{} },
		},
		readCmd: {
			linkKind:      func() interface{} { return &pb.ReadLinkRequest{} },
			namespaceKind: func() interface{} { return &pb.ReadNamespaceRequest{} },
			aliasKind:     func() interface{} { return &pb.ReadAliasRequest{} },
			targetKind:    func() interface{} { return &pb.ReadTargetRequest{} },
		},
		updateCmd: {
			linkKind:      func() interface{} { return &pb.UpdateLinkRequest{} },
			namespaceKind: func() interface{} { return &pb.UpdateNamespaceRequest{} },
			aliasKind:     func() interface{} { return &pb.UpdateAliasRequest{} },
			targetKind:    func() interface{} { return &pb.UpdateTargetRequest{} },
		},
		deleteCmd: {
			linkKind:      func() interface{} { return &pb.DeleteLinkRequest{} },
			namespaceKind: func() interface{} { return &pb.DeleteNamespaceRequest{} },
			aliasKind:     func() interface{} { return &pb.DeleteAliasRequest{} },
			targetKind:    func() interface{} { return &pb.DeleteTargetRequest{} },
		},
	}
}

func communicate(cmd *cobra.Command, _ []string) error {
	entity, err := entities.new(cmd)
	if err != nil {
		return err
	}
	if dry, _ := cmd.Flags().GetBool("dry-run"); dry {
		cmd.Printf("%T would be sent with data: ", entity)
		if cmd.Flag("output").Value.String() == jsonFormat {
			return json.NewEncoder(cmd.OutOrStdout()).Encode(entity)
		}
		return json.NewEncoder(cmd.OutOrStdout()).Encode(entity)
	}
	response, err := call(cnf.Union.GRPCConfig, entity)
	if err != nil {
		cmd.Println(err)
		return nil
	}
	if cmd.Flag("output").Value.String() == jsonFormat {
		return json.NewEncoder(cmd.OutOrStdout()).Encode(response)
	}
	return yaml.NewEncoder(cmd.OutOrStdout()).Encode(response)
}

func printSchemas(cmd *cobra.Command, _ []string) error {
	var (
		target   *cobra.Command
		builders map[kind]builder
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
	for k, b := range builders {
		_ = yaml.NewEncoder(cmd.OutOrStdout()).Encode(schema{Kind: k, Payload: convert(b())})
		cmd.Println()
	}
	return nil
}

type builder func() interface{}

type kind string

type schema struct {
	Kind    kind                   `yaml:"kind"`
	Payload map[string]interface{} `yaml:"payload"`
}

type factory map[*cobra.Command]map[kind]builder

func (f factory) new(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("filename").Value.String())
	if err != nil {
		return nil, err
	}
	build, found := f[cmd][data.Kind]
	if !found {
		return nil, errors.Errorf("unknown payload type %q", data.Kind)
	}
	entity := build()
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
	defer func() { _ = conn.Close() }()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		kit.Concat(middleware.AuthScheme, " ", string(cnf.Token)))
	switch request := entity.(type) {

	case *pb.CreateLinkRequest:
		return pb.NewLinkClient(conn).Create(ctx, request)
	case *pb.CreateNamespaceRequest:
		return pb.NewNamespaceClient(conn).Create(ctx, request)
	case *pb.CreateAliasRequest:
		return pb.NewAliasClient(conn).Create(ctx, request)
	case *pb.CreateTargetRequest:
		return pb.NewTargetClient(conn).Create(ctx, request)

	case *pb.ReadLinkRequest:
		return pb.NewLinkClient(conn).Read(ctx, request)
	case *pb.ReadNamespaceRequest:
		return pb.NewNamespaceClient(conn).Read(ctx, request)
	case *pb.ReadAliasRequest:
		return pb.NewAliasClient(conn).Read(ctx, request)
	case *pb.ReadTargetRequest:
		return pb.NewTargetClient(conn).Read(ctx, request)

	case *pb.UpdateLinkRequest:
		return pb.NewLinkClient(conn).Update(ctx, request)
	case *pb.UpdateNamespaceRequest:
		return pb.NewNamespaceClient(conn).Update(ctx, request)
	case *pb.UpdateAliasRequest:
		return pb.NewAliasClient(conn).Update(ctx, request)
	case *pb.UpdateTargetRequest:
		return pb.NewTargetClient(conn).Update(ctx, request)

	case *pb.DeleteLinkRequest:
		return pb.NewLinkClient(conn).Delete(ctx, request)
	case *pb.DeleteNamespaceRequest:
		return pb.NewNamespaceClient(conn).Delete(ctx, request)
	case *pb.DeleteAliasRequest:
		return pb.NewAliasClient(conn).Delete(ctx, request)
	case *pb.DeleteTargetRequest:
		return pb.NewTargetClient(conn).Delete(ctx, request)

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
