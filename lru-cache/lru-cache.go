package lrucache

import (
	"cache/models"
	"fmt"
)

type LRUCache struct {
	queue   *models.Queue
	hashmap map[int]*models.Node
}

func (c *LRUCache) Init(capacity int) {
	c.queue = &models.Queue{}
	c.queue.SetCapacity(capacity)
	c.hashmap = make(map[int]*models.Node, capacity)
}

func (c *LRUCache) Get(key int) interface{} {
	val, ok := c.hashmap[key]
	if !ok {
		return nil
	}
	c.queue.MoveToFront(val)
	return val.Data.Value
}

func (c *LRUCache) Put(key int, item interface{}) {
	// Check if key already exists
	// if exists map the node and replace the value
	existingNode, ok := c.hashmap[key]
	if ok {
		existingNode.Data.Value = item
		return
	}
	node, ok := c.queue.Insert(key, item)
	if !ok {
		deletedNode, ok := c.queue.Delete()
		if ok {
			node, _ := c.queue.Insert(key, item)
			c.hashmap[key] = node
			delete(c.hashmap, deletedNode.Data.Key)
		}
		return
	}
	c.hashmap[key] = node
}

func (c *LRUCache) Print() {
	node := c.queue.GetFront()
	for node != nil {
		fmt.Printf("Key: %v, Value: %v\n", node.Data.Key, node.Data.Value)
		node = node.Next
	}
}
