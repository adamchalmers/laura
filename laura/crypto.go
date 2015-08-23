package laura

import (
	"strings"
)

// Encrypt a diary.
func encrypt(plaintext string, key string) string {
	return crypto(plaintext, key, 1)
}

// Decrypt a diary.
func decrypt(cryptext string, key string) string {
	return crypto(cryptext, key, -1)
}

// Generic symmetric crypto.
func crypto(input string, key string, sign int) string {
	n := len(input)
	output := make([]uint8, n)
	for i := 0; i < n; i++ {
		// k is an integer taken from successive characters of the password.
		k := int(key[i%len(key)])
		// Add or subtract k from the input to produce the output.
		output[i] = input[i] + uint8(k*sign)
	}
	return string(output)
}

func checkPassword(cryptext string, key string) bool {
	return len(cryptext) <= 1 || strings.Contains(decrypt(cryptext, key), DELIMITER)
}
