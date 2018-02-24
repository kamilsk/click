package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// RootCmd is the entry point.
var RootCmd = &cobra.Command{Short: "Click!"}

func init() {
	RootCmd.AddCommand(migrateCmd, runCmd)
}

func must(actions ...func() error) {
	for _, action := range actions {
		if err := action(); err != nil {
			panic(err)
		}
	}
}

func asBool(value fmt.Stringer) bool {
	is, _ := strconv.ParseBool(value.String())
	return is
}

func asDuration(value fmt.Stringer) time.Duration {
	duration, _ := time.ParseDuration(value.String())
	return duration
}

func asInt(value fmt.Stringer) int {
	integer, _ := strconv.ParseInt(value.String(), 10, 0)
	return int(integer)
}
