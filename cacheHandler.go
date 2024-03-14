// cacheHandler.go
package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Data        *OutputBlockData
	LastUpdated time.Time
}

type Cache struct {
	mu     sync.RWMutex
	items  map[string]CacheItem
	expiry time.Duration
}

func NewCache(expiry time.Duration) *Cache {
	cache := &Cache{
		items:  make(map[string]CacheItem),
		expiry: expiry,
	}
	go cache.startCleanupRoutine()
	return cache
}

func (c *Cache) Get(key string) (*OutputBlockData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || time.Since(item.LastUpdated) > c.expiry {
		return nil, false
	}
	return item.Data, true
}

func (c *Cache) Set(key string, data *OutputBlockData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem{
		Data:        data,
		LastUpdated: time.Now(),
	}
}

func (c *Cache) startCleanupRoutine() {
	for {
		time.Sleep(c.expiry)
		c.Cleanup()
	}
}

func (c *Cache) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, item := range c.items {
		if time.Since(item.LastUpdated) > c.expiry {
			delete(c.items, key)
		}
	}
}

func generateCacheKey(inputData InputBlockData) string {
	return inputData.BlockType + "-" + inputData.BlocksData.Template
}

func transformToOutputBlockDataWithCache(inputData InputBlockData, cache *Cache) OutputBlockData {
	cacheKey := generateCacheKey(inputData)
	if data, found := cache.Get(cacheKey); found {
		return *data
	}

	outputData := transformToOutputBlockData(inputData)
	cache.Set(cacheKey, &outputData)

	return outputData
}
