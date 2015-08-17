package laura

import (
	// "github.com/adamchalmers/laura/filesys"
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
