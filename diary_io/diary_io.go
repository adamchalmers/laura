package diary_io

import (
	"io/ioutil"
	"os"
)

const (
	DIARY_ROOT           = "/Users/adam/Documents/laura/"
	DIARY_PERMISSION     = 0731
	DIARY_FILE_EXTENSION = ".diary"
)

// Maps diary names to their location in the filesystem.
func DiaryPath(diaryName string) string {
	return DIARY_ROOT + diaryName + DIARY_FILE_EXTENSION
}

func MakeLauraDir() error {
	return os.MkdirAll(DIARY_ROOT, DIARY_PERMISSION)
}

func MakeDiary(name string) (*os.File, error) {
	return os.Create(DiaryPath(name))
}

func FindDiaryNames() []string {
	files, err := ioutil.ReadDir(DIARY_ROOT)
	if err != nil {
		panic(err)
	}
	output := make([]string, 0)
	for _, f := range files {
		name := f.Name()
		// Strip off the .diary file extension
		output = append(output, name[:len(name)-len(DIARY_FILE_EXTENSION)])
	}
	return output
}

func ReadDiary(diaryName string) ([]byte, error) {
	return ioutil.ReadFile(DiaryPath(diaryName))
}

func DeleteDiary(name string) error {
	return os.Remove(DiaryPath(name))
}

func AddTo(name string, newText []byte) {
	ioutil.WriteFile(DiaryPath(name), newText, DIARY_PERMISSION)

}
