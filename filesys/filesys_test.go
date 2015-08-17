package filesys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNames(t *testing.T) {
	fs := NewFakeFS()
	assert.Empty(t, fs.GetNames())

	diaries := []string{"My Diary", "My Other Diary", "Leland's Diary"}
	for _, name := range diaries {
		fs.MakeDiary(name)
	}
	assert.Equal(t, diaries, fs.GetNames())
}

func TestContents(t *testing.T) {
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
