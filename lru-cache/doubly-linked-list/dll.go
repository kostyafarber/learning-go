package dll

import (
	"bytes"
	"fmt"
)

type Node struct {
	Val  rune
	Next *Node
	Prev *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%c", n.Val)
}

func (n Node) Links() string {
	return fmt.Sprintf("%v <- %c -> %v", n.Prev, n.Val, n.Next)
}

type DoublyLinkedList struct {
	Head  *Node
	Tail  *Node
	Nodes []*Node
}

func (ll DoublyLinkedList) String() string {
	var buf bytes.Buffer
	for node := ll.Head; node != nil; node = node.Next {
		fmt.Fprintf(&buf, "<- %c ->", node.Val)
	}
	return buf.String()
}

func (ll *DoublyLinkedList) deleteNodeFromArray(i int) {
	ll.Nodes = append(ll.Nodes[:i], ll.Nodes[i+1:]...)
}

func (ll *DoublyLinkedList) InsertHead(val rune) error {
	if ll.Head == nil {
		node := CreateNode(val, nil, nil)
		ll.Nodes = append(ll.Nodes, node)
		ll.Head = ll.Nodes[0]
		ll.Tail = ll.Nodes[0]
		return nil
	}

	newHead := CreateNode(val, ll.Head, nil)
	ll.Head.Prev = newHead
	ll.Nodes = append([]*Node{newHead}, ll.Nodes...)
	ll.Head = ll.Nodes[0]
	ll.Tail = ll.Nodes[ll.Len()-1]
	return nil
}

func (ll *DoublyLinkedList) DeleteTail() error {
	ll.Tail.Prev.Next = nil
	ll.Tail = ll.Tail.Prev
	ll.deleteNodeFromArray(len(ll.Nodes) - 1)
	return nil
}

func (ll *DoublyLinkedList) DeleteNode(deleteNodeVal rune) error {
	for i, node := range ll.Nodes {
		if node.Val == deleteNodeVal {
			ll.deleteNodeFromArray(i)
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}
	}
	return nil
}

func (ll *DoublyLinkedList) Len() int {
	return len(ll.Nodes)
}

func CreateNode(val rune, next *Node, prev *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
		Prev: prev,
	}
}
