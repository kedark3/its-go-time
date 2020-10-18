package linkedlistsum

import "fmt"

// LeetCode problem https://leetcode.com/problems/add-two-numbers/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	// initialize dummy node
	p := &ListNode{}
	// initialize result to point to dummy node
	res := p
	// while neither of these are not nil
	for l1 != nil || l2 != nil {
		// sum is initially set to carry
		s := carry
		// if l1 is not nil add its val to sum
		// move l1 to l1.next
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		// if l2 is not nil add its val to sum
		// move l2 to l2.next
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		// Our dummy node.Next will pooint to new Node
		// That contains Val : s % 10
		res.Next = &ListNode{Val: s % 10}
		// extract carry from sum
		carry = s / 10
		// increment res to point to newly added node
		res = res.Next
	}
	// Create new node to hold any leftover carry
	if carry > 0 {
		res.Next = &ListNode{Val: carry}
	}
	return p.Next

}

func lenOfLinkedList(l *ListNode) int {
	currentNode := l
	var counter int
	for {
		counter++
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
		if currentNode.Next == nil {
			fmt.Printf("%+v ", currentNode.Val)
			counter++
			break
		}
	}
	return counter
}
