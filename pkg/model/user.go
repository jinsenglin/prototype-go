package model

import (
	"encoding/json"
	"sync"
)

type User struct {
	Id    int
	Name  string
	Slice []string       // for demo of deep copy
	Ptr   *int           // for demo of deep copy
	Map   map[string]int // for demo of deep copy
}

func (this *User) Copy() (that User) {
	// deep copy
	b, _ := json.Marshal(this)
	json.Unmarshal(b, &that)
	return
}

type Users struct {
	Items [9]User
	mux   sync.Mutex
}

func (this *Users) List() (items []User) {
	this.mux.Lock()
	items = this.Items[:]
	this.mux.Unlock()

	return
}

func (this *Users) Get(idx int) (item User) {
	this.mux.Lock()
	item = this.Items[idx]
	this.mux.Unlock()

	return
}

func (this *Users) Create(idx int, id int, name string) {
	this.mux.Lock()
	this.Items[idx].Id = id
	this.Items[idx].Name = name
	this.mux.Unlock()
}

func (this *Users) Update(idx int, name string) {
	this.mux.Lock()
	this.Items[idx].Name = name
	this.mux.Unlock()
}

func (this *Users) Delete(idx int) {
	this.mux.Lock()
	this.Items[idx].Id = 0
	this.Items[idx].Name = ""
	this.mux.Unlock()
}
