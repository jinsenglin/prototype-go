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

// Package model ...
package model

import (
	"encoding/json"
	"sync"
)

// User ...
type User struct {
	Id    int
	Name  string
	Slice []string       // for demo of deep copy
	Ptr   *int           // for demo of deep copy
	Map   map[string]int // for demo of deep copy
}

// Copy ...
func (this *User) Copy() (that User) {
	// deep copy
	b, _ := json.Marshal(this)
	json.Unmarshal(b, &that)
	return
}

// Users ...
type Users struct {
	Items [9]User
	mux   sync.Mutex
}

// List ...
func (this *Users) List() (items []User) {
	this.mux.Lock()
	items = this.Items[:]
	this.mux.Unlock()

	return
}

// Get ...
func (this *Users) Get(idx int) (item User) {
	this.mux.Lock()
	item = this.Items[idx]
	this.mux.Unlock()

	return
}

// Create ...
func (this *Users) Create(idx int, id int, name string) {
	this.mux.Lock()
	this.Items[idx].Id = id
	this.Items[idx].Name = name
	this.mux.Unlock()
}

// Update ...
func (this *Users) Update(idx int, name string) {
	this.mux.Lock()
	this.Items[idx].Name = name
	this.mux.Unlock()
}

// Delete ...
func (this *Users) Delete(idx int) {
	this.mux.Lock()
	this.Items[idx].Id = 0
	this.Items[idx].Name = ""
	this.mux.Unlock()
}
