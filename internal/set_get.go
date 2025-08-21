package internal

import "time"

func (db *DB) Set(key, value string, ttl time.Duration) {
	db.mu.Lock()
	db.data[key] = value
	if ttl > 0 {
		db.ttl[key] = time.Now().Add(ttl)
	}
	db.mu.Unlock()
}

func (db *DB) Get(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if expiry, ok := db.ttl[key]; ok {
		if time.Now().After(expiry) {
			delete(db.data, key)
			delete(db.ttl, key)
			return "nil\n"
		}
	}

	if val, ok := db.data[key]; ok {
		return val
	}
	return "nil\n"
}
