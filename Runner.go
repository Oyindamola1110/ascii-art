package main

import (
	"strings"
	"fmt"
)

func Runner(input string, result []string)  {
	if input == "\\n" {
		fmt.Println()
		return
	}
	words := strings.Split( input,"\\n")
	for _, word := range words {
		if word == ""{
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				startLine := (int(char)-32)* 9 + 1 
				fmt.Print(result[startLine+i])			
			}
			//fmt.Println()
		}
	}
}