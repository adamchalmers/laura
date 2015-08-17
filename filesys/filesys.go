package filesys

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type FileSys interface {
	GetNames() []string
	ReadDiary(string) (string, error)
	MakeDiary(string) error
	DeleteDiary(string) error
	AddTo(string, string)
}

type realFS struct {
	DIARY_ROOT           string
	DIARY_PERMISSION     os.FileMode
	DIARY_FILE_EXTENSION string
}

func NewFS() *realFS {
	username, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fs := new(realFS)
	fs.DIARY_ROOT = fmt.Sprintf("%v/Documents/laura/", username.HomeDir)
	fs.DIARY_PERMISSION = 0731
	fs.DIARY_FILE_EXTENSION = ".diary"
	err = os.MkdirAll(fs.DIARY_ROOT, fs.DIARY_PERMISSION)
	if err != nil {
		log.Fatal(err)
	}
	return fs
}

func (fs *realFS) diaryPath(name string) string {
	return fs.DIARY_ROOT + name + fs.DIARY_FILE_EXTENSION
}

func (fs *realFS) MakeDiary(name string) error {
	_, err := os.Create(fs.diaryPath(name))
	return err
}

func (fs *realFS) GetNames() []string {
	files, err := ioutil.ReadDir(fs.DIARY_ROOT)
	if err != nil {
		log.Fatal(err)
	}
	output := make([]string, 0)
	for _, f := range files {
		name := f.Name()
		// Strip off the .diary file extension
		output = append(output, name[:len(name)-len(fs.DIARY_FILE_EXTENSION)])
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

func (fs *realFS) AddTo(name string, newText string) {
	ioutil.WriteFile(fs.diaryPath(name), []byte(newText), fs.DIARY_PERMISSION)

}
