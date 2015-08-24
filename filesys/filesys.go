/*
 * Interface for storing and using Laura journals in a real or fake filesystem.
 */

package filesys

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

const (
	DIARY_PERMISSION     = 0731
	DIARY_FILE_EXTENSION = ".diary"
	CHECKSTRING          = "!!!"
)

type FileSys interface {
	// List all diaries.
	Names() []string
	// Get the contents of a diary.
	ReadDiary(name string) (contents string, err error)
	// Make a new empty diary.
	MakeDiary(name string) error
	// Delete a diary.
	DeleteDiary(name string) error
	// Append a string to the end of a diary.
	AddTo(name string, text string) error
}

// An implementation of FileSys that stores data in a real OS filesystem.
type realFS struct {
	rootDir string
}

// Factory. Guarantees correct initialization of the private realFS struct.
func NewFS(rootDir string) *realFS {
	fs := realFS{rootDir}
	err := os.MkdirAll(fs.rootDir, DIARY_PERMISSION)
	handle(err)
	return &fs
}

// Helper function to return the absolute path to a diary.
func (fs *realFS) diaryPath(name string) string {
	return fs.rootDir + name + DIARY_FILE_EXTENSION
}

func (fs *realFS) MakeDiary(name string) error {
	_, err := os.Create(fs.diaryPath(name))
	if err != nil {
		return err
	}
	return nil
}

func (fs *realFS) Names() []string {
	files, err := ioutil.ReadDir(fs.rootDir)
	handle(err)
	output := make([]string, 0)
	for _, f := range files {
		name := f.Name()
		// Strip off the .diary file extension
		output = append(output, name[:len(name)-len(DIARY_FILE_EXTENSION)])
	}
	return output
}

func (fs *realFS) ReadDiary(diaryName string) (string, error) {
	text, err := ioutil.ReadFile(fs.diaryPath(diaryName))
	return string(text), err
}

func (fs *realFS) DeleteDiary(name string) error {
	return os.Remove(fs.diaryPath(name))
}

func (fs *realFS) AddTo(name string, newText string) error {
	return ioutil.WriteFile(fs.diaryPath(name), []byte(newText), DIARY_PERMISSION)
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func rootDir() string {
	username, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	rootDir := fmt.Sprintf("%v/Documents/laura/", username.HomeDir)
	return rootDir
}
