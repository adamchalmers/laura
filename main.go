package main

/*

Principles:
 * Append-only files. Everything was true at one point. No need to delete anything.
*/

import (
	"github.com/adamchalmers/laura/cmds"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "laura"}
	for _, cmd := range cmds.MakeCommands() {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
