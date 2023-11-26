package mod_02_cache

import (
	"sync"
	"time"
)

type CacheStorage interface {
	Get(key string) interface{}
	Set(key string, value interface{}, expire time.Duration)
	Delete(key string)
}

type CacheRecord struct {
	value  interface{}
	expire time.Time
}

type MemoryCacheStorage struct {
	storage map[string]CacheRecord
	mutex   *sync.Mutex
}

func New() CacheStorage {
	return &MemoryCacheStorage{
		storage: make(map[string]CacheRecord),
		mutex:   new(sync.Mutex),
	}
}

func (cs *MemoryCacheStorage) Get(key string) interface{} {
	cs.mutex.Lock()

	value, ok := cs.storage[key]

	if !ok {
		return nil
	}

	// Expired
	if time.Now().After(value.expire) {
		delete(cs.storage, key)

		return nil
	}

	cs.mutex.Unlock()

	return value.value
}

func (cs *MemoryCacheStorage) Set(key string, value interface{}, expire time.Duration) {
	cs.storage[key] = CacheRecord{
		value:  value,
		expire: time.Now().Add(expire),
	}
}

func (cs *MemoryCacheStorage) Delete(key string) {
	delete(cs.storage, key)
}
