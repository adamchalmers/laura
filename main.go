package main

import (
	"github.com/adamchalmers/laura/filesys"
	"github.com/adamchalmers/laura/laura"
	"github.com/spf13/cobra"
	"time"
)

func main() {
	lfs := filesys.NewFS()
	var rootCmd = &cobra.Command{Use: "laura"}
	for _, cmd := range laura.MakeCommands(lfs, time.Now()) {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
