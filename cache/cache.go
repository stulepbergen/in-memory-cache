package cache

type dictionary map[string]interface{}

type Cache struct {
	cacheMemory dictionary
}

func New() Cache {
	return Cache{
		cacheMemory: dictionary{},
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.cacheMemory[key] = value
}

func (c Cache) Get(key string) interface{} {
	value, contains := c.cacheMemory[key]

	if contains {
		return value
	}

	return "Such key doesn't exist yet!"
}

func (c Cache) Delete(key string) {
	_, contains := c.cacheMemory[key]

	if !contains {
		panic("Can't delete not existing key!")
	}

	delete(c.cacheMemory, key)
}
