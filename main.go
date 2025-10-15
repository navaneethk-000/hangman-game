package main

import "fmt"

func getSecretWord(string) string {
	return "navaneeth"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
