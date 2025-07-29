// Last Third Element
// // Create sample linked list: 10 -> 20 -> 30 -> 40 -> 50
// Third element from the end is: 30

package main

import "fmt"

type ThirdLinkedList struct {
	Val  int
	Next *ThirdLinkedList
}

func Len(l *ThirdLinkedList) int {
	lenCount := 0
	head := l
	for head != nil {

		lenCount += 1
		head = head.Next
	}
	return lenCount
}

func getTheNthNodeFromLast(l *ThirdLinkedList, n int) *ThirdLinkedList {
	head := l
	lenOftheLinkedList := Len(l)
	i := 0
	for head != nil {
		fmt.Println(head.Val, i, lenOftheLinkedList-n-1)
		if i > lenOftheLinkedList-n-1 {
			return head
		}
		i++
		head = head.Next
	}

	return nil

}

func main() {
	Node1 := &ThirdLinkedList{Val: 10}
	Node2 := &ThirdLinkedList{Val: 20}
	Node3 := &ThirdLinkedList{Val: 30}
	Node4 := &ThirdLinkedList{Val: 40}
	Node5 := &ThirdLinkedList{Val: 50}

	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node4
	Node4.Next = Node5

	lenOftheLinkedList := Len(Node1)
	nl := getTheNthNodeFromLast(Node1, 3)
	fmt.Println("The len of the Linked List is ", lenOftheLinkedList)
	fmt.Println("\n 3rd Node from the last is  ", nl.Val)

}
