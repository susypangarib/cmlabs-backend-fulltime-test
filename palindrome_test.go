package cmlabs_backend_fulltime

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func palindrome(word string) bool {
	idx := len(word) - 1
	chars := strings.Split(word, "")
	for i := 0; i < len(word)/2; i++ {
		if chars[i] != chars[idx-i] {
			return false
		}
	}
	return true
}

func TestPalindrome(t *testing.T) {
	assert.Equal(t, false, palindrome("Anna"))
	assert.Equal(t, true, palindrome("anna"))
	assert.Equal(t, false, palindrome("Aibohphobia"))
	assert.Equal(t, true, palindrome("aibohphobia"))
	assert.Equal(t, true, palindrome("SAIPPUAKIVIKAUPPIAS"))
	assert.Equal(t, true, palindrome("civic"))
	assert.Equal(t, false, palindrome("Civic"))
	assert.Equal(t, false, palindrome("My gym"))
	assert.Equal(t, true, palindrome("mygym"))
	assert.Equal(t, false, palindrome("No lemon, no melon"))
	assert.Equal(t, true, palindrome("nolemon,nomelon"))
}
