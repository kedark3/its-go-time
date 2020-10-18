package main

import "fmt"

type node struct {
	key      int
	next     *node
	previous *node
}

type list struct {
	name    string
	head    *node
	tail    *node
	current *node
}

func newList(name string) *list {
	return &list{
		name: name,
	}
}

func (l *list) addItem(item int) error {
	n := &node{
		key: item,
	}
	if l.head == nil {
		l.head = n
	} else {
		currentNode := l.tail
		currentNode.next = n
		n.previous = l.tail
	}
	l.tail = n
	return nil
}

func (l *list) startIterating() *node {
	l.current = l.head
	return l.current
}

func (l *list) nextItem() *node {
	l.current = l.current.next
	return l.current
}

func (l *list) previousItem() *node {
	l.current = l.current.previous
	return l.current
}

func (l *list) printItems() error {
	currentNode := l.head
	if currentNode == nil {
		fmt.Println("List is empty")
	}
	//	fmt.Printf("%+v\n -> ", currentNode.key)

	for {
		fmt.Printf("%+v -> ", currentNode.key)
		currentNode = currentNode.next
		if currentNode.next == nil {
			fmt.Printf("%+v ", currentNode.key)
			break
		}
	}

	return nil
}

func main() {
	myList := newList("MyList")
	fmt.Println("Created new LinkedList")
	fmt.Println()
	fmt.Println("Adding new items to LinkedList")
	for _, val := range []int{1, 2, 4, 8, 16, 32} {
		myList.addItem(val)
	}
	fmt.Println()

	myList.printItems()
	fmt.Println()
	fmt.Printf("%+v", myList.startIterating().key)
	fmt.Printf("%+v", myList.nextItem().key)
	fmt.Printf("%+v", myList.nextItem().key)
	fmt.Printf("%+v", myList.nextItem().key)
	fmt.Printf("%+v", myList.nextItem().key)
	fmt.Printf("%+v", myList.nextItem().key)

	fmt.Printf("%+v", myList.previousItem().key)
	fmt.Printf("%+v", myList.previousItem().key)
	fmt.Printf("%+v", myList.previousItem().key)
	fmt.Printf("%+v", myList.previousItem().key)
	fmt.Printf("%+v", myList.previousItem().key)

}
