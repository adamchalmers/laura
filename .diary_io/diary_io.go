package diary_io

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

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
