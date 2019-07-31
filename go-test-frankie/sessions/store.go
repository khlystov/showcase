package sessions

import "sync"

// Store is used to save unique Session keys 
// across all calls for the life of the running process
var Store = &store{Cache: make(map[string]bool)}

type store struct {
	mu    sync.Mutex
	Cache map[string]bool
}

func (store *store) Set(key string) {
	exist := store.Check(key)

	if !exist {
		store.mu.Lock()
		store.Cache[key] = true
		store.mu.Unlock()
	}
}

func (store *store) Check(key string) bool {
	store.mu.Lock()
	_, ok := store.Cache[key]
	store.mu.Unlock()

	return ok
}
