package cache

// Pop returns and removes an item from cache if it was found and not expired.
func (c *cache) Pop(k string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, found := c.items[k]
	if !found || v.Expired() {
		return nil, false
	}

	// remove key from cache
	delete(c.items, k)
	return v.Object, true
}
