package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/kamilsk/click/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Click! server via gRPC",
	}
)

func init() {
	var (
		cnf  = config.GRPCConfig{}
		file = ""
		v    = viper.New()
	)
	{
		must(
			func() error { return v.BindEnv("click_token") },
		)
		v.SetDefault("click_token", "")
	}
	{
		flags := controlCmd.PersistentFlags()
		flags.StringVarP(&file, "file", "f", file, "entity source (default is stdin)")
		flags.StringVarP(&cnf.Interface, "host", "H", "127.0.0.1:8092", "gRPC host")
		flags.DurationVarP(&cnf.Timeout, "timeout", "t", time.Second, "connection timeout")
		flags.StringVarP((*string)(&cnf.Token), "token", "", v.GetString("click_token"), "user access token")
	}
	controlCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "Create some kind",
			RunE: func(cmd *cobra.Command, args []string) error {

				var (
					err error
					out struct {
						Kind    string                 `yaml:"kind"`
						Payload map[string]interface{} `yaml:"payload"`
					}
					raw []byte
					src io.Reader = os.Stdin
				)
				if file != "" {
					if src, err = os.Open(file); err != nil {
						return err
					}
				}
				if raw, err = ioutil.ReadAll(src); err != nil {
					return err
				}
				if err = yaml.Unmarshal(raw, &out); err != nil {
					return err
				}

				log.Println("`ctl create` was called", out)
				return nil
			},
		},
		&cobra.Command{
			Use:   "get",
			Short: "Get some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl get` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "update",
			Short: "Update some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl update` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "delete",
			Short: "Delete some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl delete` was called")
				return nil
			},
		},
	)
}