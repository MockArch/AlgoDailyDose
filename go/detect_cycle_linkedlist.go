// How to find if a linked list has a loop?

// node1(1) -> node2(2) -> node3(3) -> node4(4) -> node5(5)
//                                      ^                      |
//                                      |______________________|

// Nodes 1 through 5 are connected in sequence.

// But node5 points back to node3, creating a cycle (loop).

// So if you follow the .Next pointers starting at node1, after reaching node5,
// you jump back to node3, and this cycle repeats infinitely.

// Step-by-step explanation of cycle detection
// Two pointers start at head:

// slow moves 1 step at a time

// fast moves 2 steps at a time

// If no cycle, fast reaches the end (nil)

// If cycle exists, slow and fast will eventually meet inside the cycle

// To find cycle start:

// Move slow back to head

// Move both slow and fast 1 step at a time

// Where they meet again is the start of the cycle

package main

import "fmt"

type NodeList struct {
	Val  int
	Next *NodeList
}

func hasCycle(head *NodeList) bool {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		if slow == fast {
			return true

		}
	}

	return false
}

func main() {

	Node1 := &NodeList{Val: 1}
	Node2 := &NodeList{Val: 2}
	Node3 := &NodeList{Val: 3}
	Node4 := &NodeList{Val: 4}
	Node5 := &NodeList{Val: 5}

	Node1.Next = Node2
	Node2.Next = Node3
	Node3.Next = Node4
	Node4.Next = Node5

	// This creates a loop: 3 -> 4 -> 5 -> 3 -> ...
	Node5.Next = Node3

	if hasCycle(Node1) {
		fmt.Println("Cycle detected!")
	}

}
