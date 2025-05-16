package ttlmap

import (
	"sync"
	"time"
)

// item is a struct that holds the value and the last access time
type item struct {
	value      interface{}
	lastAccess int64
}

// TTLMap is a map that holds items with a time-to-live (TTL)
type TTLMap struct {
	m map[string]*item
	// For safe access to the map
	mu sync.Mutex
}

// New creates a new TTLMap with the given size and maxTTL
func New(size int, maxTTL int) (m *TTLMap) {
	// map is created with the given length
	m = &TTLMap{m: make(map[string]*item, size)}

	// this goroutine will clean up the map from old items
	go func() {
		// You can adjust this ticker to be more or less frequent
		for now := range time.Tick(time.Second) {
			m.mu.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.mu.Unlock()
		}
	}()

	return
}
