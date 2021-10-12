package main

import "fmt"


type Node struct {
	property int 
	nextNode *Node
	previousNode *Node 

}


type LinkedList struct {
	headNode *Node
}


func (linkedList *LinkedList) NodeBetweenValues(firstProp,secondProp int) *Node{
	var node *Node
	var nodeWithVal *Node

	for node = linkedList.headNode;node != nil;

	node = node.nextNode{
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProp && node.nextNode.property == secondProp{
				nodeWithVal = node
				break
			}
		}
	}
	fmt.Println("node with val between is ---->",nodeWithVal)
	return nodeWithVal
}