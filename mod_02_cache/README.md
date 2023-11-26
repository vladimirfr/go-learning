# cache

## In-memory Cache Storage

Allows to implement in-memory caching.

### Basic Methods

 - `Set(key string, value interface{}, expire time.Duration)` - writing the value `value` to the cache using the key `key`, `expire` time to expire;
 - `Get(key string)` - reading value from cache using key `key`;
 - `Delete(key)` - deleting a value from the cache using the key `key`.

###  Import Module

```
go get -u github.com/vladimirfr/go-learning/02-cache
```

### Example 

```
func main() {
    cache := cache.New()

    cache.Set("userId", 42, time.Second * 60)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
}
```