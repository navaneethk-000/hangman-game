package main

import (
	"bufio"
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

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
