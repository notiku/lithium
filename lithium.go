package lithium

import (
	"container/list"

	"github.com/notiku/lithium/rules"
	"github.com/notiku/lithium/rules/lru"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Invalidate(key string) bool
	InvalidateContaining(substring string)
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
