package main

/*

Principles:
 * Append-only files. Everything was true at one point. No need to delete anything.
*/

import (
	"github.com/adamchalmers/laura/cmds"
	"github.com/adamchalmers/laura/filesys"
	"github.com/spf13/cobra"
)

func main() {
	lfs := filesys.NewFS()
	var rootCmd = &cobra.Command{Use: "laura"}
	for _, cmd := range cmds.MakeCommands(lfs) {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
