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

func (u1 *User) Copy() (u2 User) {
	// deep copy
	b, _ := json.Marshal(u1)
	json.Unmarshal(b, &u2)
	return
}

type Users struct {
	Items [9]User
	mux   sync.Mutex
}

func (data *Users) List() []User {
	data.mux.Lock()
	us := data.Items[:]
	data.mux.Unlock()

	return us
}

func (data *Users) Get(idx int) (u User) {
	data.mux.Lock()
	u = data.Items[idx]
	data.mux.Unlock()

	return
}

func (data *Users) Create(idx int, id int, name string) {
	data.mux.Lock()
	data.Items[idx].Id = id
	data.Items[idx].Name = name
	data.mux.Unlock()
}

func (data *Users) Update(idx int, name string) {
	data.mux.Lock()
	data.Items[idx].Name = name
	data.mux.Unlock()
}

func (data *Users) Delete(idx int) {
	data.mux.Lock()
	data.Items[idx].Id = 0
	data.Items[idx].Name = ""
	data.mux.Unlock()
}
