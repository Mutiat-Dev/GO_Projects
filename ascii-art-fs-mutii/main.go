package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	text := ""
	banner := "standard"

	if len(os.Args) == 2 {
		text = os.Args[1]
	} else if len(os.Args) == 3 {
		text = os.Args[1]
		banner = os.Args[2]
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		fmt.Println("invalid banner")
		return
	}

	file := banner + ".txt"

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	words := strings.Split(text, "\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {

				asciiIndex := 9 * int(char-32) + 1
				fmt.Print(lines[asciiIndex+i])
			}
			fmt.Println()
		}
	}
}
