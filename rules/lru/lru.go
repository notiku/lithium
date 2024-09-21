package lru

import (
	"container/list"
	"sync"
)

// Cache is a least-recently-used cache.
type Cache struct {
	MaxSize int
	Cache   map[string]*list.Element
	Ll      *list.List
	mu      sync.Mutex
}

type cacheItem struct {
	key string
	val interface{}
}

// Get retrieves a value from the cache.
// If the key does not exist, it will return nil and false.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.Cache[key]; ok {
		c.Ll.MoveToFront(ele)
		return ele.Value.(*list.Element).Value, true
	}
	return nil, false
}

// Set adds a value to the cache.
// If the cache is full, it will remove the least-recently-used item.
// If the key already exists, it will update the value and move the item to the front.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.Cache[key]; ok {
		c.Ll.MoveToFront(ele)
		ele.Value.(*cacheItem).val = value
		return
	}

	item := &cacheItem{key: key, val: value}
	ele := c.Ll.PushFront(item)
	c.Cache[key] = ele

	if c.Ll.Len() > c.MaxSize {
		c.removeOldest()
	}
}

func (c *Cache) removeOldest() {
	ele := c.Ll.Back()
	if ele != nil {
		c.Ll.Remove(ele)
		item := ele.Value.(*cacheItem)
		delete(c.Cache, item.key)
	}
}

// Invalidate removes a value from the cache.
// If the key does not exist, it will return false.
func (c *Cache) Invalidate(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.Cache[key]; ok {
		c.Ll.Remove(ele)
		delete(c.Cache, key)
		return true
	}
	return false
}

// InvalidateContaining removes all values from the cache that contain the provided substring.
func (c *Cache) InvalidateContaining(substring string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.Cache {
		if contains(k, substring) {
			c.Invalidate(k)
		}
	}
}

// GetStats returns the number of items in the cache and the maximum size.
func (c *Cache) GetStats() (int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Ll.Len(), c.MaxSize
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}
