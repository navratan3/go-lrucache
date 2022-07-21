package cache

type Cache struct {
	capacity int
	strategy Strategy
}

func (c *Cache) Init() {
	c.strategy.Init(c.capacity)
}

func (c *Cache) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Cache) SetCapacity(capacity int) {
	c.capacity = capacity
}

func (c *Cache) GetCapcity() int {
	return c.capacity
}

func (c *Cache) Get(key int) interface{} {
	return c.strategy.Get(key)
}

func (c *Cache) Put(key int, item interface{}) {
	c.strategy.Put(key, item)
}

func (c *Cache) Print() {
	c.strategy.Print()
}
