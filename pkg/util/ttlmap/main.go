package ttlmap

import (
	"time"
)

// Put adds a new item to the map or updates the existing one
func (m *TTLMap) Put(k string, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, ok := m.m[k]
	if !ok {
		it = &item{
			value: v,
		}
	}
	it.value = v
	it.lastAccess = time.Now().Unix()
	m.m[k] = it
}

// Get returns the value of the given key if it exists
func (m *TTLMap) Get(k string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if it, ok := m.m[k]; ok {
		it.lastAccess = time.Now().Unix()
		return it.value, true
	}

	return nil, false
}

// Delete removes the item from the map
func (m *TTLMap) Delete(k string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.m[k]; ok {
		delete(m.m, k)
	}
}
