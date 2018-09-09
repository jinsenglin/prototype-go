/*
Additional Resources
- https://golang.org/pkg/io/#Copy
- https://golang.org/pkg/io/ioutil/#WriteFile
- https://golang.org/pkg/bufio/#Writer.Write
- https://golang.org/pkg/bufio/#Writer.WriteString
- https://golang.org/pkg/os/#File.Write
- https://golang.org/pkg/os/#File.WriteString
- https://godoc.org/golang.org/x/time/rate
- https://studygolang.com/articles/10148
- https://github.com/fujiwara/shapeio
- https://gobyexample.com/writing-files
- https://segmentfault.com/a/1190000011680507
- https://gist.github.com/ebraminio/576fdfdff425bf3335b51a191a65dbdb
- https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html
- http://mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket/
*/

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func _ioutil_WriteFile(from string, to string) {
	if fileFrom, err := os.Open(from); err != nil {
		log.Println(err)
	} else {
		defer fileFrom.Close()
		if data, err := ioutil.ReadFile(from); err != nil {
			log.Println(err)
		} else {
			if err := ioutil.WriteFile(to, data, 0644); err != nil {
				log.Println(err)
			} else {
				log.Println("copied")
			}
		}
	}
}

func _io_Copy(from string, to string) {
	if fileFrom, err := os.Open(from); err != nil {
		log.Println(err)
	} else {
		defer fileFrom.Close()
		if fileTo, err := os.Create(to); err != nil {
			log.Println(err)
		} else {
			defer fileTo.Close()
			if _, err := io.Copy(fileTo, fileFrom); err != nil {
				log.Println(err)
			} else {
				if err := fileFrom.Sync(); err != nil {
					log.Println(err)
				} else {
					log.Println("copied")
				}
			}
		}
	}
}

func cp(from string, to string) {
	log.Printf("coping file %v to %v\n", from, to)
	// _io_Copy(from, to)
	_ioutil_WriteFile(from, to)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Need exactly two arguments.")
	} else {
		cp(os.Args[1], os.Args[2])
	}
}
