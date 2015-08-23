package main

import (
	"fmt"
	"github.com/adamchalmers/laura/filesys"
	"github.com/adamchalmers/laura/laura"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"time"
)

func main() {
	lfs := filesys.NewFS()
	var rootCmd = &cobra.Command{Use: "laura"}

	var Key = func() string {
		fmt.Printf("Please enter your password: ")
		bytes, err := terminal.ReadPassword(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		return string(bytes)
	}

	for _, cmd := range laura.MakeCommands(lfs, time.Now(), Key) {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.Execute()

}
