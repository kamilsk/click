package cmd

import (
	"log"

	"github.com/kamilsk/click/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Click! server via gRPC",
	}
)

func init() {
	var (
		flags = controlCmd.PersistentFlags()
		cnf   = config.GRPCConfig{}
		v     = viper.New()
	)
	{
		must(
			func() error { return v.BindEnv("click_token") },
		)
		v.SetDefault("click_token", "")
	}
	{
		flags.StringVarP((*string)(&cnf.Token), "token", "t", v.GetString("click_token"), "user access token")
	}
	controlCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "Create some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl create` was called")
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
