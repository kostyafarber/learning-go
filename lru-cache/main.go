package main

import (
	"fmt"
	lrucache "lru-cache/lru-cache"
)

func main() {
	cache := lrucache.CreateCache(5)

	ops := []rune{'A', 'B', 'C', 'D', 'F', 'C', 'Z', 'F'}
	fmt.Printf("Simulating operations on LRU (Least Recently Used Cache) of: \n%c\n", ops)

	for _, c := range ops {
		_, ok := cache.Get(c)

		if ok == -1 {
			fmt.Printf("Cache miss, adding %c to the cache\n", c)
			fmt.Printf("Adding %c to the cache\n", c)
			fmt.Printf("After adding %c, cache is: %v\n\n", c, cache)
		} else {
			fmt.Printf("%c in the cache, retrieving it\n", c)
			fmt.Printf("cache is: %v\n\n", cache)
		}
	}

}
