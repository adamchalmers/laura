package laura

import (
	"github.com/adamchalmers/laura/filesys"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCrypto(t *testing.T) {
	msg := "Hello, world!"
	key := "Password"
	assert.Equal(t, msg, encrypt(decrypt(msg, key), key))
	assert.Equal(t, msg, decrypt(encrypt(msg, key), key))
}

func TestAddToDiary(t *testing.T) {
	now := time.Now()
	key := "Password"
	oldText := encrypt("previous entry text.", key)
	newText := "Hello, world."
	contents := addToDiaryText(oldText, newText, now, key)
	assert.Contains(t, decrypt(contents, key), timestamp(now))
	assert.Contains(t, decrypt(contents, key), newText)
}

func TestEnforceArgs(t *testing.T) {
	assert.True(t, missingArgs([]string{"X"}, "abc def"))
	assert.False(t, missingArgs([]string{"X", "Y"}, "abc def"))
}

func TestAddtoCommand(t *testing.T) {
	lfs := filesys.NewFakeFS()
	now := time.Now()
	key := "Password"
	var keyFn = func() string {
		return key
	}
	cmds := MakeCommands(lfs, now, keyFn)

	// Make a new diary
	cmds["new"].SetArgs([]string{"journal"})
	cmds["new"].Execute()
	assert.Contains(t, lfs.Names(), "journal")

	// Write to it
	cmds["addto"].SetArgs([]string{"journal", "Hello world"})
	cmds["addto"].Execute()
	contents, err := lfs.ReadDiary("journal")
	assert.Nil(t, err)
	assert.Contains(t, decrypt(contents, key), "Hello world")
	assert.Contains(t, decrypt(contents, key), timestamp(now))

	// Write a second entry
	cmds["addto"].SetArgs([]string{"journal", "Applecake"})
	cmds["addto"].Execute()
	contents, err = lfs.ReadDiary("journal")
	assert.Nil(t, err)
	assert.Contains(t, decrypt(contents, key), "Applecake")
	assert.Contains(t, decrypt(contents, key), timestamp(now))
}
