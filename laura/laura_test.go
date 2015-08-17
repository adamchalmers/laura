package laura

import (
	"github.com/adamchalmers/laura/filesys"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCrypto(t *testing.T) {
	msg := "Hello, world!"
	assert.Equal(t, msg, encrypt(decrypt(msg)))
	assert.Equal(t, msg, decrypt(encrypt(msg)))
}

func TestAddToDiary(t *testing.T) {
	now := time.Now()
	oldText := encrypt("previous entry text.")
	newText := "Hello, world."
	contents := addToDiaryText(oldText, newText, now)
	assert.Contains(t, contents, encrypt(timestamp(now)))
	assert.Contains(t, contents, encrypt(newText))
}

func TestEnforceArgs(t *testing.T) {
	assert.True(t, missingArgs([]string{"X"}, "abc def"))
	assert.False(t, missingArgs([]string{"X", "Y"}, "abc def"))
}

func TestAddtoCommand(t *testing.T) {
	lfs := filesys.NewFakeFS()
	now := time.Now()
	cmds := MakeCommands(lfs, now)

	// Make a new diary
	cmds["new"].SetArgs([]string{"journal"})
	cmds["new"].Execute()
	assert.Contains(t, lfs.Names(), "journal")

	// Write to it
	msgs := []string{"Hello, world.", "Apple cake"}
	for _, msg := range msgs {
		cmds["addto"].SetArgs([]string{"journal", msg})
		cmds["addto"].Execute()
		contents, err := lfs.ReadDiary("journal")
		assert.Nil(t, err)
		assert.Contains(t, contents, encrypt(msg))
		assert.Contains(t, contents, encrypt(timestamp(now)))
	}
}
