package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	DIARY_ROOT           = "/Users/adam/Documents/laura/"
	DIARY_PERMISSION     = 0731
	DIARY_FILE_EXTENSION = ".diary"
)

func makeCmd(name string, desc string, argNames string, fn func(*cobra.Command, []string)) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", name, argNames),
		Short: desc,
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			if argNames != "" {
				enforceArgs(args, argNames)
			}
			fn(cmd, args)
		},
	}
}

var CmdNew = makeCmd("new", "Makes a new diary", "diaryName", func(cmd *cobra.Command, args []string) {
	err := os.MkdirAll(DIARY_ROOT, DIARY_PERMISSION)
	dealWith(err)
	_, err = os.Create(diaryPath(args[0]))
	dealWith(err)
})

var CmdList = makeCmd("list", "Lists all diaries", "", func(cmd *cobra.Command, args []string) {
	files, err := ioutil.ReadDir(DIARY_ROOT)
	for _, f := range files {
		name := f.Name()
		fmt.Println(name[:len(name)-len(DIARY_FILE_EXTENSION)]) // Strip off the .diary file extension
	}
	dealWith(err)
})

var CmdAddto = makeCmd("addto", "Adds text to a diary", "diaryName newEntryText", func(cmd *cobra.Command, args []string) {
	diaryName, newEntryText := args[0], args[1]

	year, month, day := time.Now().Date()
	timestamp := fmt.Sprintf("%d %s %d\n", day, month, year)

	text, err := ioutil.ReadFile(diaryPath(diaryName))
	dealWith(err)
	text = decrypt(text)
	newPlaintext := fmt.Sprintf("%s%s---\n%s\n\n", text, timestamp, newEntryText)
	newText := encrypt([]byte(newPlaintext))

	ioutil.WriteFile(diaryPath(diaryName), newText, DIARY_PERMISSION)
})

var CmdRead = makeCmd("read", "Displays contents of a diary", "diaryName", func(cmd *cobra.Command, args []string) {
	diaryName := args[0]

	bytes, err := ioutil.ReadFile(diaryPath(diaryName))
	if err != nil {
		fmt.Printf("Couldn't find diary '%s'\n", diaryName)
		return
	}
	text := string(decrypt(bytes))

	fmt.Println(text)
})

var CmdDelete = makeCmd("delete", "Delete a diary", "diaryName", func(cmd *cobra.Command, args []string) {
	diaryName := args[0]

	err := os.Remove(diaryPath(diaryName))
	if err != nil {
		fmt.Printf("No diary named %s exists\n", diaryName)
	}
})

func dealWith(err error) {
	if err != nil {
		panic(err)
	}
}

// Maps diary names to their location in the filesystem.
func diaryPath(diaryName string) string {
	return DIARY_ROOT + diaryName + DIARY_FILE_EXTENSION
}

// Exits the program with an error message if len(args) < len(expected args).
func enforceArgs(actual []string, expected string) {
	if len(actual) < len(strings.Split(expected, " ")) {
		fmt.Printf("Expected arguments [%s]\n", expected)
		os.Exit(1)
	}
}

// Encrypt a diary.
func encrypt(plaintext []byte) []byte {
	output := []byte("")
	for _, char := range plaintext {
		n := char + 1
		output = append(output, n)
	}
	return output
}

// Decrypt a diary.
func decrypt(cryptext []byte) []byte {
	output := []byte("")
	for _, char := range cryptext {
		n := char - 1
		output = append(output, n)
	}
	return output
}
