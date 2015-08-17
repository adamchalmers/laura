/*
 * App logic for Laura.
 */

package laura

import (
	"fmt"
	"github.com/adamchalmers/laura/filesys"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"time"
)

/*
 * Make all the command objects to be used with the Cobra library.
 * Takes in a FileSys object (dependency injection) for testing purposes.
 */
func MakeCommands(lfs filesys.FileSys) []*cobra.Command {

	makeCmd := func(name string, desc string, argNames string, fn func(*cobra.Command, []string)) *cobra.Command {
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
		err := lfs.MakeDiary(args[0])
		if err != nil {
			fmt.Printf(err.Error())
		}
	})

	var CmdList = makeCmd("list", "Lists all diaries", "", func(cmd *cobra.Command, args []string) {
		for _, name := range lfs.GetNames() {
			fmt.Println(name)
		}
	})

	var CmdAddto = makeCmd("addto", "Adds text to a diary", "diaryName newEntryText", func(cmd *cobra.Command, args []string) {
		diaryName, newEntryText := args[0], args[1]

		text, err := lfs.ReadDiary(diaryName)
		dealWith(err)

		newText := addToDiaryText(text, newEntryText, time.Now())

		lfs.AddTo(diaryName, newText)
	})

	var CmdRead = makeCmd("read", "Displays contents of a diary", "diaryName", func(cmd *cobra.Command, args []string) {
		diaryName := args[0]

		bytes, err := lfs.ReadDiary(diaryName)
		if err != nil {
			fmt.Printf("Couldn't find diary '%s'\n", diaryName)
			return
		}
		text := readDiary(bytes)

		fmt.Println(text)
	})

	var CmdDelete = makeCmd("delete", "Delete a diary", "diaryName", func(cmd *cobra.Command, args []string) {
		diaryName := args[0]

		err := lfs.DeleteDiary(diaryName)
		if err != nil {
			fmt.Printf("No diary named %s exists\n", diaryName)
		}
	})

	return []*cobra.Command{CmdNew, CmdList, CmdAddto, CmdRead, CmdDelete}
}

func addToDiaryText(cryptext string, newEntryText string, t time.Time) string {
	text := decrypt(cryptext)
	year, month, day := t.Date()
	timestamp := fmt.Sprintf("%d %s %d\n", day, month, year)
	newPlaintext := fmt.Sprintf("%s%s---\n%s\n\n", text, timestamp, newEntryText)
	newText := encrypt(newPlaintext)
	return newText
}

func readDiary(bytes string) string {
	return string(decrypt(bytes))
}

// Exits the program with an error message if len(args) < len(expected args).
func enforceArgs(actual []string, expected string) {
	if len(actual) < len(strings.Split(expected, " ")) {
		fmt.Printf("Expected arguments [%s]\n", expected)
		os.Exit(1)
	}
}

// Encrypt a diary.
func encrypt(plaintext string) string {
	output := ""
	for _, char := range plaintext {
		n := char + 1
		output += string(n)
	}
	return output
}

// Decrypt a diary.
func decrypt(cryptext string) string {
	output := ""
	for _, char := range cryptext {
		n := char - 1
		output += string(n)
	}
	return output
}

func dealWith(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
