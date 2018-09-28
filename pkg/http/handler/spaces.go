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

package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func spaceID(path string) (id int) {
	re, _ := regexp.Compile("[1-9]")
	id, _ = strconv.Atoi(re.FindString(path))
	return
}

// SpacesAPIHandler ...
func SpacesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if re, _ := regexp.Compile("^/spaces/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			id := spaceID(r.URL.Path)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><frameset rows='*,50px'><frame src='/channels/%d/chats'><frame src='/chats/new?ch_id=%d'></frameset></html>", id, id)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		log.Println(r.URL.Path)
		http.NotFound(w, r)
	}
}
