
package interfaces

// Cacheable is an interface for managing cache entries.
type CacheAble interface {
	// Get returns the value for the given key.
	Get(key string) ([]byte, error)
	// Set sets the value for the given key.
	Set(key string, value interface{}, ttl int) error
	// SetWithExpireAt set key value and update expire using unix timestamp
}