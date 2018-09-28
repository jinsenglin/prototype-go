//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

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
