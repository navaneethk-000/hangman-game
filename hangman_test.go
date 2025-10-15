package main

import (
	"strings"
	"testing"
)

func TestSecretWordNoCapitals(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if secretWord != strings.ToLower(secretWord) {
		t.Errorf("Should not get words with capital letters. Got %s", secretWord)
	}
}

func TestSecretWordLength(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWordLength := len(getSecretWord(wordList))
	if secretWordLength < 6 {
		t.Errorf("Shoulf not get words of length less than 6. Got %v", secretWordLength)
	}
}
