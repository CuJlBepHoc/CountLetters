package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.ToLower(text)
	text = strings.Replace(text, "\n", "", -1)
	
	counts := make(map[rune]int)
	totalChars := 0

	for _, char := range text {
		if unicode.IsLetter(char) {
			counts[char]++
			totalChars++
		}
	}

	for char, count := range counts {
		percentage := float64(count) / float64(totalChars) * 100
		fmt.Printf("%c - %d %.2f%%\n", char, count, percentage)
	}
}
