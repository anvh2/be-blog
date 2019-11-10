package cache

// Cache ...
type Cache interface {
	Get(key string) interface{}
	Put(key string, value interface{}) interface{}
	Delete(key string)
	Size() int
}
