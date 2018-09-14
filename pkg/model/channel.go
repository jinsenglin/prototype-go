package model

import "sync"

type Channel struct {
	Slice []Chat
}

type Channels struct {
	Items [9]Channel
	mux   sync.Mutex
}
