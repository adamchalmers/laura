package main

/*

Principles:
 * Append-only files. Everything was true at one point. No need to delete anything.
*/

import (
	"github.com/adamchalmers/laura/cmds"
	"github.com/adamchalmers/laura/diary_io"
	"github.com/spf13/cobra"
)

func main() {
	lfs := new(diary_io.RealFS)
	var rootCmd = &cobra.Command{Use: "laura"}
	for _, cmd := range cmds.MakeCommands(lfs) {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
