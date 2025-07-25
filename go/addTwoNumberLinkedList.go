// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]
// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) Append(val int) {
	newNode := &ListNode{Val: val, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) TraverseAndPrint() {
	fmt.Println("Traversing the linked list:")
	current := l.Head

	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func arrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}

func listToArray(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		digit := sum % 10

		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummyHead.Next
}

func main() {
	fmt.Println("--- Main Function Examples ---")

	list1Head := arrayToList([]int{2, 4, 3})
	list2Head := arrayToList([]int{5, 6, 4})

	fmt.Print("List 1 (342): ")
	tempList1 := &LinkedList{Head: list1Head}
	tempList1.TraverseAndPrint()

	fmt.Print("List 2 (465): ")
	tempList2 := &LinkedList{Head: list2Head}
	tempList2.TraverseAndPrint()

	sumListHead := addTwoNumbers(list1Head, list2Head)
	fmt.Print("Sum List (Expected: 7 -> 0 -> 8 -> nil): ")
	if sumListHead != nil {
		tempSumList := &LinkedList{Head: sumListHead}
		tempSumList.TraverseAndPrint()
	} else {
		fmt.Println("nil (addTwoNumbers not yet implemented)")
	}
	fmt.Printf("Array from Sum List: %v\n\n", listToArray(sumListHead))

	list3Head := arrayToList([]int{0})
	list4Head := arrayToList([]int{0})

	fmt.Print("List 3 (0): ")
	tempList3 := &LinkedList{Head: list3Head}
	tempList3.TraverseAndPrint()

	fmt.Print("List 4 (0): ")
	tempList4 := &LinkedList{Head: list4Head}
	tempList4.TraverseAndPrint()

	sumList2Head := addTwoNumbers(list3Head, list4Head)
	fmt.Print("Sum List (Expected: 0 -> nil): ")
	if sumList2Head != nil {
		tempSumList2 := &LinkedList{Head: sumList2Head}
		tempSumList2.TraverseAndPrint()
	} else {
		fmt.Println("nil (Error or unexpected state)")
	}
	fmt.Printf("Array from Sum List: %v\n\n", listToArray(sumList2Head))

	list5Head := arrayToList([]int{9, 9, 9, 9, 9, 9, 9})
	list6Head := arrayToList([]int{9, 9, 9, 9})

	fmt.Print("List 5 (9999999): ")
	tempList5 := &LinkedList{Head: list5Head}
	tempList5.TraverseAndPrint()

	fmt.Print("List 6 (9999): ")
	tempList6 := &LinkedList{Head: list6Head}
	tempList6.TraverseAndPrint()

	sumList3Head := addTwoNumbers(list5Head, list6Head)
	fmt.Print("Sum List (Expected: 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1 -> nil): ")
	if sumList3Head != nil {
		tempSumList3 := &LinkedList{Head: sumList3Head}
		tempSumList3.TraverseAndPrint()
	} else {
		fmt.Println("nil (Error or unexpected state)")
	}
	fmt.Printf("Array from Sum List: %v\n\n", listToArray(sumList3Head))

	fmt.Println("--- End Main Function Examples ---")
}
