package filesys

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFakeNames(t *testing.T) {
	fs := NewFakeFS()
	assert.Empty(t, fs.Names())

	diaries := []string{"My Diary", "My Other Diary", "Leland's Diary"}
	for _, name := range diaries {
		fs.MakeDiary(name)
	}
	assert.Equal(t, diaries, fs.Names())
}

func TestFakeContents(t *testing.T) {
	fs := NewFakeFS()
	diaries := map[string]string{
		"A": "This is the first diary.",
		"B": "Here's another diary.",
		"C": "The last diary?",
	}
	for name := range diaries {

		// Make an empty diary, check it's actually empty.
		fs.MakeDiary(name)
		startContents, err := fs.ReadDiary(name)
		assert.Nil(t, err)
		assert.Equal(t, "", startContents)

		// Add text to the diary.
		fs.AddTo(name, diaries[name])
		actual, err := fs.ReadDiary(name)
		assert.Nil(t, err)
		assert.Equal(t, diaries[name], actual)
	}
}

func TestRealNames(t *testing.T) {
	dir := rootDir() + "test/"
	fs := NewFS(dir)
	assert.Empty(t, fs.Names())
	diaries := []string{"MyDiary", "MyOtherDiary", "Leland'sDiary"}
	for _, name := range diaries {
		fs.MakeDiary(name)
	}
	for _, name := range diaries {
		assert.Contains(t, fs.Names(), name)
		fs.DeleteDiary(name)
		assert.NotContains(t, fs.Names(), name)
	}
	assert.Empty(t, fs.Names())
	os.Remove(dir)
}
