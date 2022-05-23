package cache

import "time"

var maxTime = time.Unix(1<<63-62135596801, 999999999)

type Cache struct {
	items map[string]CacheItem
}

type CacheItem struct {
	value       string
	timeExpired time.Time
}

func NewCache() Cache {
	return Cache{items: map[string]CacheItem{}}
}

func (c *Cache) Get(key string) (string, bool) {
	item, exists := c.items[key]
	if !exists {
		return "", false
	}

	if time.Now().After(item.timeExpired) {
		delete(c.items, key)
		return "", false
	}

	return item.value, true
}

func (c *Cache) Put(key, value string) {
	c.items[key] = CacheItem{value: value, timeExpired: maxTime}
}

func (c *Cache) Keys() []string {
	ret := make([]string, 0, len(c.items))
	for k, v := range c.items {
		if time.Now().After(v.timeExpired) {
			delete(c.items, k)
		} else {
			ret = append(ret, k)
		}
	}

	return ret
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = CacheItem{value: value, timeExpired: deadline}
}
