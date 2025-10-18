package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func createDictFile(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")

	if err != nil {
		fmt.Println("Couldn't create temp file.")
	}

	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "Parrot", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}

	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestSecretWordNoPunctuation(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)

	if err != nil {
		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}

	secretWord := getSecretWord(wordList)

	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFile([]string{"lion", "pen", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}

	secretWord := getSecretWord(wordList)

	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestCorrectGuess(t *testing.T) {

	secretWord := "elephant"
	userInput := 'e'
	state := NewGame(secretWord)

	newState := checkGuess(state, byte(userInput))

	expected := Game{
		secretWord:       secretWord,
		chancesRemaining: 7,
		guessedLetters:   append(state.guessedLetters, byte(userInput)),
		correctGuesses:   append(state.correctGuesses, byte(userInput)),
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if newState.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Remaining chances modified")
	}
	if string(newState.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Expected %q but got %q", expected.guessedLetters, newState.guessedLetters)
	}
	if string(newState.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Expected %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

}

func TestCorrectGuess2(t *testing.T) {
	secretWord := "elephant"
	userInput := 'e'
	state := Game{
		secretWord:       secretWord,
		chancesRemaining: 7,
		guessedLetters:   []byte{'l', 'b', 'z'},
		correctGuesses:   []byte{'l'},
	}

	newState := checkGuess(state, byte(userInput))

	expected := Game{
		secretWord:       state.secretWord,
		chancesRemaining: state.chancesRemaining,
		guessedLetters:   append(state.guessedLetters, byte(userInput)),
		correctGuesses:   append(state.correctGuesses, byte(userInput)),
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if newState.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Remaining chances modified")
	}
	if string(newState.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Expected %q but got %q", expected.guessedLetters, newState.guessedLetters)
	}
	if string(newState.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Expected %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}
}

func TestWrongGuess(t *testing.T) {
	secretWord := "elephant"
	userInput := 'z'
	currentState := Game{
		secretWord:       secretWord,
		chancesRemaining: 7,
		guessedLetters:   []byte{'e'},
		correctGuesses:   []byte{'e'},
	}

	newState := checkGuess(currentState, byte(userInput))
	expected := Game{
		secretWord:       secretWord,
		chancesRemaining: 6,
		guessedLetters:   []byte{'e', 'z'},
		correctGuesses:   []byte{'e'},
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if newState.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Remaining chances modified")
	}
	if string(newState.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Expected %q but got %q", expected.guessedLetters, newState.guessedLetters)
	}
	if string(newState.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Expected %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

}

func TestAlreadyGuessed(t *testing.T) {
	secretWord := "elephant"
	currentState := Game{
		secretWord:       secretWord,
		chancesRemaining: 6,
		guessedLetters:   []byte{'e', 'z'},
		correctGuesses:   []byte{'e'},
	}
	userInput := 'z'
	newState := checkGuess(currentState, byte(userInput))
	expected := Game{
		secretWord:       secretWord,
		chancesRemaining: 6,
		guessedLetters:   []byte{'e', 'z'},
		correctGuesses:   []byte{'e'},
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if newState.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Remaining chances modified")
	}
	if string(newState.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Expected %q but got %q", expected.guessedLetters, newState.guessedLetters)
	}
	if string(newState.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Expected %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

}
