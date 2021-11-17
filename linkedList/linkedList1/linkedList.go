package main

import "fmt"

type node struct {
	val  int
	next *node
}

type llist struct {
	head    *node
	tail    *node
	_length int
}

func newList() *llist {
	return &llist{}
}

func newNode(val int) *node {
	return &node{
		val:  val,
		next: nil,
	}
}

func (l *llist) len() int {
	return l._length
}
func (l *llist) addItem(val int) error {
	if l.head == nil {
		l.head = newNode(val)
		l._length++
	}
	if l.tail == nil {
		l.tail = l.head
		return nil
	} else if l.head != nil {
		currNode := l.head
		for ; currNode.next != nil; currNode = currNode.next {
		}
		currNode.next = newNode(val)
		currNode = currNode.next
		l.tail = currNode
		l._length++
		return nil
	}

	return fmt.Errorf("couldn't add %d to the list", val)
}

func (l *llist) deleteItem(pos int) error {
	if pos > l.len() {
		return fmt.Errorf("list is not as long as given position %d", pos)
	}
	if l.head == l.tail && pos == 0 {
		fmt.Printf("\nSingle element list - Deleting %d\n", l.head.val)
		l.head, l.tail = nil, nil
		l._length--
		return nil
	} else if pos == 0 {
		prevNode := l.head
		l.head = l.head.next
		prevNode.next = nil
		fmt.Printf("\nDeleted %d\n", prevNode.val)
		l._length--
		return nil
	}
	prevNode, currNode := l.head, l.head
	for i := 1; i < pos; i++ {
		prevNode = currNode
		currNode = currNode.next
	}
	fmt.Printf("\nDeleting %d at %d\n", currNode.val, pos)
	prevNode.next = currNode.next
	currNode = nil
	l._length--
	return nil
}

func (l *llist) print() {
	currNode := l.head
	for ; currNode != nil; currNode = currNode.next {
		fmt.Printf("%3d", currNode.val)
	}
	fmt.Println()

}

func (l *llist) insert(val, pos int) error {
	if pos > l.len() {
		return fmt.Errorf("asking to insert at %d in the list of size %d, please provide valid position", pos, l.len())
	}

	if pos == 0 {
		newNode := newNode(val)
		newNode.next = l.head
		l.head = newNode
	} else if pos == l.len() {
		newNode := newNode(val)
		l.tail.next = newNode
		l.tail = newNode
	} else {
		currNode, prevNode := l.head, l.head
		for i := 1; i < pos; i++ {
			prevNode = currNode
			currNode = currNode.next
		}
		newNode := newNode(val)
		fmt.Printf("\nInserting %d at %d\n", val, pos)
		prevNode.next = newNode
		newNode.next = currNode
	}
	return nil

}

func main() {
	l := newList()
	for i := 1; i <= 10; i++ {
		err := l.addItem(i)
		if err != nil {
			fmt.Println("Some error occured", err)
			return
		}
	}
	fmt.Printf("Linked list has %d items.\n", l.len())
	l.print()
	// l.deleteItem(0)
	// l.print()
	// l.deleteItem(5)
	// l.print()
	// l.deleteItem(7)
	// l.print()
	// fmt.Printf("Linked list has %d items.\n", l.len())
	// l.insert(11, 0)
	// l.print()
	// l.insert(22, l.len())
	// l.print()
	l.insert(22, 5)
	l.print()
}
