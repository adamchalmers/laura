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
	assert.Contains(t, contents, encrypt(timestamp(now), key))
	assert.Contains(t, contents, encrypt(newText, key))
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
	msgs := []string{"Hello, world.", "Apple cake"}
	for _, msg := range msgs {
		cmds["addto"].SetArgs([]string{"journal", msg})
		cmds["addto"].Execute()
		contents, err := lfs.ReadDiary("journal")
		assert.Nil(t, err)
		assert.Contains(t, contents, encrypt(msg, key))
		assert.Contains(t, contents, encrypt(timestamp(now), key))
	}
}
