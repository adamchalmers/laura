package cmds

import (
	"fmt"
	"github.com/adamchalmers/laura/diary_io"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

func MakeCommands() []*cobra.Command {
	return []*cobra.Command{CmdNew, CmdList, CmdAddto, CmdRead, CmdDelete}
}

func addToDiaryText(cryptext []byte, newEntryText string, t time.Time) []byte {
	text := decrypt(cryptext)
	year, month, day := t.Date()
	timestamp := fmt.Sprintf("%d %s %d\n", day, month, year)
	newPlaintext := fmt.Sprintf("%s%s---\n%s\n\n", text, timestamp, newEntryText)
	newText := encrypt([]byte(newPlaintext))
	return newText
}

func readDiary(bytes []byte) string {
	return string(decrypt(bytes))
}

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

var CmdNew = makeCmd("new", "Makes a new diary", "diaryName", func(cmd *cobra.Command, args []string) {
	err := diary_io.MakeLauraDir()
	if err != nil {
		fmt.Printf(err.Error())
	}
	_, err = diary_io.MakeDiary(args[0])
	if err != nil {
		fmt.Printf(err.Error())
	}
})

var CmdList = makeCmd("list", "Lists all diaries", "", func(cmd *cobra.Command, args []string) {
	for _, name := range diary_io.FindDiaryNames() {
		fmt.Println(name)
	}
})

var CmdAddto = makeCmd("addto", "Adds text to a diary", "diaryName newEntryText", func(cmd *cobra.Command, args []string) {
	diaryName, newEntryText := args[0], args[1]

	text, err := diary_io.ReadDiary(diaryName)
	dealWith(err)

	newText := addToDiaryText(text, newEntryText, time.Now())

	diary_io.AddTo(diaryName, newText)
})

var CmdRead = makeCmd("read", "Displays contents of a diary", "diaryName", func(cmd *cobra.Command, args []string) {
	diaryName := args[0]

	bytes, err := diary_io.ReadDiary(diaryName)
	if err != nil {
		fmt.Printf("Couldn't find diary '%s'\n", diaryName)
		return
	}
	text := readDiary(bytes)

	fmt.Println(text)
})

var CmdDelete = makeCmd("delete", "Delete a diary", "diaryName", func(cmd *cobra.Command, args []string) {
	diaryName := args[0]

	err := diary_io.DeleteDiary(diaryName)
	if err != nil {
		fmt.Printf("No diary named %s exists\n", diaryName)
	}
})

func dealWith(err error) {
	if err != nil {
		panic(err)
	}
}
