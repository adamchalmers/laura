package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

/*

Principles:
 * Append-only files. Everything was true at one point. No need to delete anything.
 *

Laura API:
 $ laura addto myjournal
 This will read stdin, encrypt it and add it to the file 'myjournal'
 $ laura read myjournal
 This will echo the plaintext from myjournal.

*/
const DIARY_ROOT = "/Users/adam/Documents/laura/"
const DIARY_PERMISSION = 0731

func main() {

	var cmdNew = &cobra.Command{
		Use:   "new",
		Short: "Makes a new diary",
		Long:  "Makes a new, empty diary.",
		Run: func(cmd *cobra.Command, args []string) {
			enforceArgs(args, "diaryName")
			err := os.MkdirAll(DIARY_ROOT, DIARY_PERMISSION)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			_, err = os.Create(diaryPath(args[0]))
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Lists all diaries",
		Long:  "Lists all diaries.",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := ioutil.ReadDir(DIARY_ROOT)
			for _, f := range files {
				name := f.Name()
				fmt.Println(name[:len(name)-6])
			}
			dealWith(err)
		},
	}

	var cmdAddto = &cobra.Command{
		Use:   "addto",
		Short: "Adds text to a diary",
		Long:  "Adds text to a diary",
		Run: func(cmd *cobra.Command, args []string) {
			enforceArgs(args, "diaryName newEntryText")
			diaryName, newEntryText := args[0], args[1]

			text, err := ioutil.ReadFile(diaryPath(diaryName))
			text = decrypt(text)
			dealWith(err)
			newText := encrypt([]byte(string(text) + newEntryText + "\n"))
			ioutil.WriteFile(diaryPath(diaryName), newText, DIARY_PERMISSION)
		},
	}

	var cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Displays contents of a diary",
		Long:  "Displays contents of a diary",
		Run: func(cmd *cobra.Command, args []string) {
			enforceArgs(args, "diaryName")
			diaryName := args[0]

			bytes, err := ioutil.ReadFile(diaryPath(diaryName))
			if err != nil {
				fmt.Printf("Couldn't find diary '%s'\n", diaryName)
				return
			}
			text := string(decrypt(bytes))

			fmt.Println(text)
		},
	}

	var rootCmd = &cobra.Command{Use: "laura"}
	rootCmd.AddCommand(cmdNew, cmdList, cmdAddto, cmdRead)
	rootCmd.Execute()

}

func dealWith(err error) {
	if err != nil {
		panic(err)
	}
}

func diaryPath(diaryName string) string {
	return DIARY_ROOT + diaryName + ".diary"
}

// Exits the program with an error message if len(args) < n
func enforceArgs(actual []string, expected string) {
	if len(actual) < len(strings.Split(expected, " ")) {
		fmt.Printf("Expected arguments [%s]\n", expected)
		os.Exit(1)
	}
}

func encrypt(plaintext []byte) []byte {
	output := []byte("")
	for _, char := range plaintext {
		n := char + 1
		output = append(output, n)
	}
	return output
}

func decrypt(cryptext []byte) []byte {
	output := []byte("")
	for _, char := range cryptext {
		n := char - 1
		output = append(output, n)
	}
	return output
}
