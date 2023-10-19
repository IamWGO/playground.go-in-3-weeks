// # M 1.7 The Robber's Language
// # Make a program that translates a string into the robber language.
// # ### Code from ChatGPT because I don't understand the logic


package main

import (
	"fmt"
	"strings"
)

func toRovarspraket(text string) string {
	vowels := "AEIOUaeiou"
	var translatedText string

	for _, char := range text {
		charStr := string(char)
		if strings.Contains(vowels, charStr) {
			translatedText += charStr
		} else if strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", charStr) {
			translatedText += charStr + "o" + charStr
		} else {
			translatedText += charStr
		}
	}

	return translatedText
}

func main() {
	inputText := "Hello"
	robberLanguageText := toRovarspraket(inputText)
	fmt.Println(robberLanguageText)

	fmt.Println("\n-- End program Medium1.7 --\n")
} 