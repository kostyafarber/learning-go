package main

import (
	"fmt"
	dll "lru-cache/doubly-linked-list"
)

type LruCache struct {
	Capacity int
	cache    map[rune]dll.Node
	dll      dll.DoublyLinkedList
}

func (lru LruCache) String() string {
	return fmt.Sprintf("%v", lru.cache)
}

func (lru LruCache) Get(val rune) {

}

func (lru LruCache) Put(val rune) {

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
	ll := CreateLinkedList()
	secondNode := ll.Nodes[1]

	fmt.Printf("The doubly linked list is: %v\n\n", ll)
	fmt.Printf("Nodes in the list are: %v\n", ll.Nodes)

	fmt.Printf("Second node in the list is: %v\n", secondNode)
	fmt.Printf("Second node links are: %s\n\n", secondNode.Links())

	fmt.Printf("head links are: %s\n", ll.Head.Links())
	fmt.Printf("tail links are: %s\n\n", ll.Tail.Links())

	fmt.Println("Deleting node: C")
	ll.DeleteNode('C')

	fmt.Println(ll)
	fmt.Println(ll.Nodes)

	fmt.Println("\nInserting new value at head: Z")
	ll.InsertHead('Z')
	fmt.Println(ll)
	fmt.Println(ll.Nodes)

	fmt.Println("\nDeleting tail node")
	ll.DeleteTail()
	fmt.Println(ll)
	fmt.Println(ll.Nodes)
}
