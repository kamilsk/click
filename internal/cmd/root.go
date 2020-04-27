package cmd

import "github.com/spf13/cobra"

// RootCmd is the entry point of command-line interface.
var RootCmd = &cobra.Command{Use: "click", Short: "Click!"}

func init() { RootCmd.AddCommand(completionCmd, controlCmd, migrateCmd, runCmd) }
