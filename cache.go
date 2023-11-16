package cache

type CacheStorage interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
}

type MemoryCacheStorage struct {
	storage map[string]interface{}
}

func New() CacheStorage {
	return &MemoryCacheStorage{
		storage: make(map[string]interface{}),
	}
}

func (cs *MemoryCacheStorage) Get(key string) interface{} {
	return cs.storage[key]
}

func (cs *MemoryCacheStorage) Set(key string, value interface{}) {
	cs.storage[key] = value
}

func (cs *MemoryCacheStorage) Delete(key string) {
	delete(cs.storage, key)
}
