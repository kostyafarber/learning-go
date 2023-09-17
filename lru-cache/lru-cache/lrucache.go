package lrucache

import (
	"fmt"
	dll "lru-cache/doubly-linked-list"
)

type LruCache struct {
	Capacity int
	cache    map[rune]dll.Node
	dll      *dll.DoublyLinkedList
}

func (lru LruCache) String() string {
	return fmt.Sprintf("%v", lru.dll.Nodes)
}

func (lru *LruCache) put(val rune) {
	lru.dll.InsertHead(val)
	node := lru.dll.Head
	lru.cache[val] = *node
}

func (lru *LruCache) makeItemMostRecent(val rune) {
	lru.dll.DeleteNode(val)
	lru.dll.InsertHead(val)
}

func (lru *LruCache) Get(val rune) (rune, int) {
	if _, ok := lru.cache[val]; !ok {
		if lru.dll.Len() == lru.Capacity {
			lru.dll.DeleteTail()
		}
		lru.put(val)
		return val, -1
	} else {
		lru.makeItemMostRecent(val)
		return val, 0
	}
}

func (lru *LruCache) Items() []*dll.Node {
	return lru.dll.Nodes
}

func CreateCache(capacity int) LruCache {
	return LruCache{
		Capacity: capacity,
		dll:      &dll.DoublyLinkedList{},
		cache:    make(map[rune]dll.Node),
	}

}

func CreateLinkedList() dll.DoublyLinkedList {
	chars := []rune{'A', 'B', 'C', 'D', 'E', 'F'}
	var ll dll.DoublyLinkedList
	head := dll.CreateNode(chars[0], nil, nil)
	ll.Head = head
	ll.Nodes = append(ll.Nodes, ll.Head)

	prevNode := head
	for _, val := range chars[1:] {
		node := dll.CreateNode(val, nil, prevNode)
		prevNode.Next = node
		prevNode = node
		ll.Nodes = append(ll.Nodes, node)
	}
	ll.Tail = prevNode

	return ll
}

func main() {
}
