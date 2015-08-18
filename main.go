package main

import (
	"fmt"
	"github.com/adamchalmers/laura/filesys"
	"github.com/adamchalmers/laura/laura"
	"github.com/spf13/cobra"
	"time"
)

func main() {
	lfs := filesys.NewFS()
	var rootCmd = &cobra.Command{Use: "laura"}

	var Key = func() string {
		fmt.Printf("Please enter your password: ")
		pw := ""
		fmt.Scanf("%s", &pw)
		return pw
	}

	for _, cmd := range laura.MakeCommands(lfs, time.Now(), Key) {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
