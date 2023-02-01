package cache

import "errors"

type Cache struct {
	cache map[string]any
}

func New() Cache {
	m := make(map[string]any)
	return Cache{m}
}

func (c *Cache) Set(key string, value any) (bool, error) {
	if _, ok := c.cache[key]; ok {
		return false, errors.New("there is already a value for this key")
	}
	c.cache[key] = value
	return true, nil
}

func (c *Cache) Delete(key string) (bool, error) {
	if _, ok := c.cache[key]; !ok {
		return false, errors.New("there is no such element")
	}
	delete(c.cache, key)
	return true, nil
}

func (c *Cache) Clear() (bool, error) {
	if len(c.cache) > 0 {
		c.cache = map[string]any{}
		return false, nil
	}
	return false, errors.New("cache is empty")
}

func (c *Cache) Get(key string) (any, error) {
	if value, ok := c.cache[key]; ok {
		return value, nil
	}
	return nil, errors.New("there is no such element")
}

func (c *Cache) GetAll() (map[string]any, error) {
	if len(c.cache) > 0 {
		return c.cache, nil
	}
	return nil, errors.New("cache is empty")
}
