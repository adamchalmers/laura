/*
 App logic for Laura.
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

const (
	DELIMITER = "---"
)

// MakeCommands returns all the Cobra commands that implement Laura's functionality.
//  lfs:   a Laura Filesystem object for interacting with diary storage.
//  t:     a Time object for the timestamps which are added to diaries.
//  keyFn: a function to extract the encryption key for the user's diaries.
func MakeCommands(lfs filesys.FileSys, t time.Time, keyFn func() string) map[string]*cobra.Command {

	makeCmd := func(name string, desc string, argNames string, fn func(*cobra.Command, []string)) *cobra.Command {
		return &cobra.Command{
			Use:   fmt.Sprintf("%s [%s]", name, argNames),
			Short: desc,
			Long:  desc,
			Run: func(cmd *cobra.Command, args []string) {
				if argNames != "" && missingArgs(args, argNames) {
					os.Exit(1)
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
		for _, name := range lfs.Names() {
			fmt.Println(name)
		}
	})

	var CmdAddto = makeCmd("addto", "Adds text to a diary", "diaryName newEntryText", func(cmd *cobra.Command, args []string) {
		diaryName, newEntryText := args[0], args[1]

		text, err := lfs.ReadDiary(diaryName)
		dealWith(err)

		newText, err := addToDiaryText(text, newEntryText, t, keyFn())
		dealWith(err)

		lfs.AddTo(diaryName, newText)
	})

	var CmdRead = makeCmd("read", "Displays contents of a diary", "diaryName", func(cmd *cobra.Command, args []string) {
		diaryName := args[0]

		bytes, err := lfs.ReadDiary(diaryName)
		if err != nil {
			fmt.Printf("Couldn't find diary '%s'\n", diaryName)
			return
		}
		key := keyFn()
		if !checkPassword(bytes, key) {
			log.Fatal("Wrong password.")
		}
		fmt.Println(string(decrypt(bytes, key)))
	})

	var CmdDelete = makeCmd("delete", "Delete a diary", "diaryName", func(cmd *cobra.Command, args []string) {
		diaryName := args[0]

		err := lfs.DeleteDiary(diaryName)
		if err != nil {
			fmt.Printf("No diary named %s exists\n", diaryName)
		}
	})

	return map[string]*cobra.Command{
		"new":    CmdNew,
		"list":   CmdList,
		"addto":  CmdAddto,
		"read":   CmdRead,
		"delete": CmdDelete,
	}
}

func addToDiaryText(cryptext string, newEntryText string, t time.Time, key string) (string, error) {
	if !checkPassword(cryptext, key) {
		return "", fmt.Errorf("Wrong password.")
	}
	old := decrypt(cryptext, key)
	ts := timestamp(t)
	return encrypt(formatEntry(old, newEntryText, ts), key), nil
}

func formatEntry(oldText string, newText string, timestamp string) string {
	return fmt.Sprintf("%s%s%s\n%s\n\n", oldText, timestamp, DELIMITER, newText)
}

func timestamp(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d %s %d\n", day, month, year)
}

// Exits the program with an error message if len(args) < len(expected args).
func missingArgs(actual []string, expected string) bool {
	if len(actual) < len(strings.Split(expected, " ")) {
		fmt.Printf("Expected arguments [%s]\n", expected)
		return true
	}
	return false
}

func dealWith(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
