This library provides an implementation for managing an in-memory cache. It allows the user to set, get, delete, and clear cache elements, with error handling for missing elements or empty caches.
Usage


import "cache"

c := cache.New()

// Set a key-value pair in the cache
c.Set("key", "value")

// Get a value by key
value, _ := c.Get("key")

// Delete a key-value pair
c.Delete("key")

// Clear the cache
c.Clear()

// Get all key-value pairs
all, _ := c.GetAll()

API Reference
func New() Cache

Returns a new cache instance.
func (c *Cache) Set(key string, value any) (bool, error)

Sets a key-value pair in the cache. Returns false and an error if there is already a value for the given key.
func (c *Cache) Delete(key string) (bool, error)

Deletes a key-value pair from the cache. Returns false and an error if the key does not exist.
func (c *Cache) Clear() (bool, error)

Clears the cache. Returns false and an error if the cache is already empty.
func (c *Cache) Get(key string) (any, error)

Gets a value from the cache by key. Returns an error if the key does not exist.
func (c *Cache) GetAll() (map[string]any, error)

Gets all key-value pairs from the cache. Returns nil and an error if the cache is empty.
