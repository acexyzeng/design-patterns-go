package LRU

import (
	"testing"
)

func TestLRU(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(1, 101)
	cache.Put(2, 102)
	cache.Put(3, 103)
	t.Log(cache.Keys)
}
