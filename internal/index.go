package internal

import (
	"sync"
	"time"
)

type DB struct {
	data map[string]string
	ttl  map[string]time.Time
	mu   sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]string),
		ttl:  make(map[string]time.Time),
	}
}
