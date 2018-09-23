package main

import (
	"strings"
	"text/scanner"

	"golang.org/x/tour/wc"
)

// Wrong answer!
// Input: "I am learning Go!"
// Expect: "{"learning":1, "Go!":1, "I":1, "am":1}"
// But return: {"I":1, "am":1, "learning":1, "Go":1, "!":1}
func implement2(s string) map[string]int {
	result := make(map[string]int)

	var ss scanner.Scanner
	ss.Init(strings.NewReader(s))
	for tok := ss.Scan(); tok != scanner.EOF; tok = ss.Scan() {
		word := ss.TokenText()

		_, ok := result[word]
		if ok {
			result[word] = result[word] + 1
		} else {
			result[word] = 1
		}
	}

	return result
}

func implement1(s string) map[string]int {
	result := make(map[string]int)

	for _, word := range strings.Split(s, " ") {
		_, ok := result[word]
		if ok {
			result[word] = result[word] + 1
		} else {
			result[word] = 1
		}
	}

	return result
}

func WordCount(s string) map[string]int {
	return implement1(s)
}

func main() {
	wc.Test(WordCount)
}
