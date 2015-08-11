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
	rootCmd.AddCommand(cmds.CmdNew, cmds.CmdList, cmds.CmdAddto, cmds.CmdRead, cmds.CmdDelete)
	rootCmd.Execute()

}
