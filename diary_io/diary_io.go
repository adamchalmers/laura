package diary_io

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type LauraFS interface {
	GetNames() []string
	ReadDiary(diaryName string) ([]byte, error)
	MakeDiary(string) (*os.File, error)
	DeleteDiary(string) error
	Init() error
	AddTo(string, []byte)
}

type RealFS struct {
	DIARY_ROOT           string
	DIARY_PERMISSION     os.FileMode
	DIARY_FILE_EXTENSION string
}

func (fs *RealFS) Init() error {
	username, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fs.DIARY_ROOT = fmt.Sprintf("%v/Documents/laura/", username.HomeDir)
	fs.DIARY_PERMISSION = 0731
	fs.DIARY_FILE_EXTENSION = ".diary"
	return os.MkdirAll(fs.DIARY_ROOT, fs.DIARY_PERMISSION)
}

func (fs *RealFS) diaryPath(name string) string {
	return fs.DIARY_ROOT + name + fs.DIARY_FILE_EXTENSION
}

func (fs *RealFS) MakeDiary(name string) (*os.File, error) {
	return os.Create(fs.diaryPath(name))
}

func (fs *RealFS) GetNames() []string {
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

func (fs *RealFS) ReadDiary(diaryName string) ([]byte, error) {
	return ioutil.ReadFile(fs.diaryPath(diaryName))
}

func (fs *RealFS) DeleteDiary(name string) error {
	return os.Remove(fs.diaryPath(name))
}

func (fs *RealFS) AddTo(name string, newText []byte) {
	ioutil.WriteFile(fs.diaryPath(name), newText, fs.DIARY_PERMISSION)

}
