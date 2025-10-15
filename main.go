package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
		if word == strings.ToLower(word) {
			allowedWords = append(allowedWords, word)
		}
	}

	randomNum := rand.Intn(len(allowedWords))

	return allowedWords[randomNum]
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
