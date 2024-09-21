package lithium

import (
	"container/list"

	"github.com/notiku/lithium/rules"
	"github.com/notiku/lithium/rules/lru"
)

type Cache interface {
	// Get retrieves a value from the cache.
	// If the key does not exist, it will return nil and false.
	Get(key string) (interface{}, bool)

	// Set adds a value to the cache.
	// If the cache is full, it will remove the least-recently-used item.
	// If the key already exists, it will update the value and move the item to the front.
	Set(key string, value interface{})

	// Invalidate removes a value from the cache.
	// If the key does not exist, it will return false.
	Invalidate(key string) bool

	// InvalidateContaining removes all values from the cache that contain the provided substring.
	InvalidateContaining(substring string)

	// GetStats returns the number of items in the cache and the maximum size.
	GetStats() (int, int)
}

// New returns a new Cache instance based on the provided strategy and parameter.
func New(s rules.Strategy, param interface{}) Cache {
	switch s {
	case rules.LRU:
		return &lru.Cache{
			MaxSize: param.(int),
			Cache:   make(map[string]*list.Element),
			Ll:      list.New(),
		}
	default:
		return nil
	}
}
