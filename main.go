package main

import (
	"fmt"
	"os"
	"strings"
)

func validator(arguments []string) bool {
	if len(arguments) < 1 {
		return false
	}
	if arguments[0] == "" {
		return false
	}
	return true
}
func readFile() (string, error) {
	content, err := os.ReadFile("./banner-files/standard.txt")

	if err != nil {
		fmt.Println("file unreadable")
		return "", err
	}
	return string(content), nil
}

 func main() {
arguments := os.Args[1:]

if !validator(arguments) {
	return
}
content, _ := readFile()

input := arguments[0]
result := strings.Split(string(content), "\n")

Runner(input, result)

}