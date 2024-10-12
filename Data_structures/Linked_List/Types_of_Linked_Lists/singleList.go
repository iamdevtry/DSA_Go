package main

import "fmt"

type Node struct {
	data  int
	pNext *Node
}

func createNode(value int) *Node {
	return &Node{data: value, pNext: nil}
}

type SingleList struct {
	pHead *Node
}

func createList() *SingleList {
	return &SingleList{pHead: nil}
}

func main() {
	node := createNode(10)
	fmt.Println("Node created with value:", node.data)
	fmt.Println("Node created with next pointer:", node.pNext)
	fmt.Println("=====================================")

	nodeList := createList()
	fmt.Println("List created with head:", nodeList.pHead)
	fmt.Println("=====================================")

	nodeList.pHead = node
	fmt.Println("List head updated with node:", nodeList.pHead)
	fmt.Println("=====================================")

	// Add a new node to the list
	newNode := createNode(20)
	fmt.Println("New node added to the list:", newNode.pNext)
	fmt.Println("=====================================")

	// Add the new node to the list
	newNode.pNext = nodeList.pHead
	nodeList.pHead = newNode
	fmt.Println("New node added to the list:", nodeList.pHead)
	fmt.Println("=====================================")

	// Traverse the list
	pTemp := nodeList.pHead
	for pTemp != nil {
		fmt.Println("Node data:", pTemp.data)
		pTemp = pTemp.pNext
	}
}
