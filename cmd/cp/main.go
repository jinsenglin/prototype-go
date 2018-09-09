/*
Additional Resources
- https://golang.org/pkg/io/#Copy
- https://golang.org/pkg/io/ioutil/#WriteFile
- https://golang.org/pkg/bufio/#Writer.Write
- https://golang.org/pkg/bufio/#Writer.WriteString
- https://golang.org/pkg/os/#File.Write
- https://golang.org/pkg/os/#File.WriteString

Rate Limit
- https://godoc.org/golang.org/x/time/rate
- https://studygolang.com/articles/10148
- https://github.com/fujiwara/shapeio
- https://medium.com/@KevinHoffman/rate-limiting-service-calls-in-go-3771c6b7c146

Compare 4 W/R Methods
- https://gobyexample.com/writing-files
- https://segmentfault.com/a/1190000011680507

2 HTTP File Transfer Methods
- r.Body https://gist.github.com/ebraminio/576fdfdff425bf3335b51a191a65dbdb
- r.FormFile https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html

TCP File Transfer
- http://mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket/
*/

package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func _os_File_Write(from string, to string) {
	if fileFrom, err := os.Open(from); err != nil {
		log.Println(err)
	} else {
		defer fileFrom.Close()
		if fileTo, err := os.Create(to); err != nil {
			log.Println(err)
		} else {
			defer fileTo.Close()
			block := make([]byte, 1) // TODO: refactor with larger block for better performance
			for {
				if n, err := fileFrom.Read(block); err != nil && err != io.EOF {
					log.Fatal(err)
				} else if 0 == n {
					if err := fileTo.Sync(); err != nil {
						log.Println(err)
					} else {
						log.Println("copied")
					}
					break
				} else {
					fileTo.Write(block)
				}
			}
		}
	}
}

func _bufio_Write(from string, to string) {
	if fileFrom, err := os.Open(from); err != nil {
		log.Println(err)
	} else {
		defer fileFrom.Close()
		if fileTo, err := os.Create(to); err != nil {
			log.Println(err)
		} else {
			defer fileTo.Close()
			fileReader := bufio.NewReader(fileFrom)
			fileWriter := bufio.NewWriter(fileTo)
			block := make([]byte, 1) // TODO: refactor with larger block for better performance
			for {
				if n, err := fileReader.Read(block); err != nil && err != io.EOF {
					log.Fatal(err)
				} else if 0 == n {
					if err := fileWriter.Flush(); err != nil {
						log.Println(err)
					} else {
						log.Println("copied")
					}
					break
				} else {
					fileWriter.Write(block)
				}
			}
		}
	}
}

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
	// _ioutil_WriteFile(from, to)
	// _bufio_Write(from, to)
	_os_File_Write(from, to)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Need exactly two arguments.")
	} else {
		cp(os.Args[1], os.Args[2])
	}
}
