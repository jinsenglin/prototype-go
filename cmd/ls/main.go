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
	"fmt"
	"log"
	"os"
)

func ls(path string) {
	fmt.Println(path, ":")

	if file, err := os.Open(path); err != nil {
		log.Println(err)
	} else {
		defer file.Close()
		if info, err := file.Readdir(-1); err != nil {
			log.Println(err)
		} else {
			for _, fi := range info {
				fmt.Println(fi.Name())
			}
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Need at least one argument.")
	} else {
		for _, arg := range os.Args[1:] {
			ls(arg)
		}
	}
}
