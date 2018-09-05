/*
Golang built-in api to list files in a directory
See https://medium.com/@manigandand/list-files-in-a-directory-using-golang-963b1df11304

Third-party api to list files in a directory
See https://github.com/spf13/afero
*/

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
		if info, err := file.Readdir(-1); err != nil {
			log.Println(err)
		} else {
			for _, fi := range info {
				fmt.Println(fi.Name())
			}
		}
		file.Close()
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
