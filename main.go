package main

import (
	"cache/cache"
	lrucache "cache/lru-cache"
	"fmt"
)

func main() {
	//Driver code for cache
	c := new(cache.Cache)
	l := new(lrucache.LRUCache)
	c.SetStrategy(l)
	c.SetCapacity(3)
	c.Init()
	var option int
	var key int
	var value string
	for {
		fmt.Printf("1. Put\n2. Get\n3. Print\n")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Enter key and value")
			fmt.Scanln(&key)
			fmt.Scanln(&value)
			c.Put(key, value)
		case 2:
			fmt.Println("Enter key")
			fmt.Scanln(&key)
			fmt.Printf("Value: %v\n", c.Get(key))
		case 3:
			fmt.Println("Printing cache")
			c.Print()
		}
	}
}
