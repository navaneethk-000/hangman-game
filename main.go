package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Game struct {
	secretWord       string
	guessedLetters   []byte
	chancesRemaining uint
	correctGuesses   []byte
}

func NewGame(secretWord string) Game {
	return Game{
		secretWord:       secretWord,
		guessedLetters:   []byte{},
		chancesRemaining: 7,
		correctGuesses:   []byte{},
	}
}

func checkGuess(state Game, userInput byte) Game {

	guess := userInput
	if state.chancesRemaining > 0 && !bytes.Contains(state.guessedLetters, []byte{guess}) {

		if strings.ContainsRune(state.secretWord, rune(guess)) {
			state.correctGuesses = append(state.correctGuesses, guess)
			state.guessedLetters = append(state.guessedLetters, guess)
		} else {
			state.guessedLetters = append(state.guessedLetters, guess)
			state.chancesRemaining--
		}
	}

	return state
}

func hasPunctuation(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return true
		}
	}
	return false
}

func getSecretWord(wordFileName string) string {
	allowedWords := []string{}
	wordFile, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Can't open file %s : %v\n", wordFileName, err)
		panic(errMessage)
	}
	defer wordFile.Close()

	scanner := bufio.NewScanner(wordFile)

	for scanner.Scan() {

		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && !hasPunctuation(word) {
			allowedWords = append(allowedWords, word)
		}
	}

	randomNum := rand.Intn(len(allowedWords))

	return allowedWords[randomNum]
}

func hasWonGame(state Game) bool {
	guessedLetters := make(map[byte]bool)

	for _, ch := range state.correctGuesses {
		guessedLetters[ch] = true
	}

	for i := 0; i < len(state.secretWord); i++ {
		if !guessedLetters[state.secretWord[i]] {
			return false
		}
	}
	return true
}

func getGuess() byte {
	fmt.Print("Guess a letter : ")
	reader := bufio.NewReader(os.Stdin)
	ch, _ := reader.ReadByte()
	reader.ReadByte()
	return ch
}

func hasLossGame(state Game) bool {
	return state.chancesRemaining == 0
}

func main() {
	secretWord := getSecretWord("/usr/share/dict/words")
	game := NewGame(secretWord)
	fmt.Println("Welcome to Hangman!")

	for game.chancesRemaining > 0 {

		if hasWonGame(game) {
			fmt.Println("You Won!")
			break
		}

		fmt.Println("Chances remaining:", game.chancesRemaining)
		fmt.Println("Guessed Letters : ", string(game.guessedLetters))

		guess := getGuess()
		game = checkGuess(game, guess)

		if hasLossGame(game) {
			fmt.Println("You Lose!")
		}
	}
}
