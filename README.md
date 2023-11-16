# cache

## In-memory Cache Storage

Allows to implement in-memory caching.

### Basic Methods

 - `Set(key string, value interface{})` - writing the value `value` to the cache using the key `key`;
 - `Get(key string)` - reading value from cache using key `key`;
 - `Delete(key)` - deleting a value from the cache using the key `key`.

###  Import Module
`go get -u github.com/vladimirfr/cache`

### Example 

`func main() {
    cache := cache.New()

    cache.Set("userId", 42)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
}`