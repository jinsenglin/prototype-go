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
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func _osFileWrite(from string, to string) {
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

func _bufioWrite(from string, to string) {
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

func _ioutilWriteFile(from string, to string) {
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

func _ioCopy(from string, to string) {
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
	// _ioCopy(from, to)
	// _ioutilWriteFile(from, to)
	// _bufioWrite(from, to)
	_osFileWrite(from, to)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Need exactly two arguments.")
	} else {
		cp(os.Args[1], os.Args[2])
	}
}
