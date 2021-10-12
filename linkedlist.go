package main

import "fmt"

//node class with property int and nextNode which points to the Node class
type Node struct{
	property int
	nextNode *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}


func (linkedList *LinkedList) AddToHead(property int){
	 var node = Node{}
	 node.property = property

	 if node.nextNode != nil{
		 node.nextNode = linkedList.headNode
	 }

	 linkedList.headNode = &node
}


func (linkedList *LinkedList) IterateList(){
	var node *Node 

	// iterated thru each node and check the head node is not empty.
	for node = linkedList.headNode; node != nil;

	node = node.nextNode{
		fmt.Println("current head node is ----->",node.property)
	}

}

func (linkedList *LinkedList) LastNode() *Node{
	var node *Node 
	var lastnode *Node 

	for node = linkedList.headNode; node != nil;

	node = node.nextNode{
		if node.nextNode == nil {
			lastnode = node
	
		}
	}
	fmt.Println("lastnode node is ---->",lastnode)
	return lastnode
}


func (linkedList *LinkedList) AddToEnd(property int){
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastnode *Node
	lastnode = linkedList.LastNode()
	if lastnode != nil{
		fmt.Println("lastnode added to end is -->",lastnode)
		lastnode.nextNode = node
	}
}

func (linkedList *LinkedList) NodeWithVal(property int) *Node{
	var node *Node 
	var nodeWithVal *Node 

	for node = linkedList.headNode; node != nil;

	node = node.nextNode{
		if  node.property == property{
			nodeWithVal = node
			break
		}
	}
	fmt.Println("node with value is ---->",nodeWithVal)
	return nodeWithVal
}

func (linkedList *LinkedList) AddAfter(nodePropToAdd, property int){

	var node = &Node{}
	node.property = property
	node.nextNode = nil 

	var nodeWithVal *Node 
	nodeWithVal = linkedList.NodeWithVal(nodePropToAdd)

	if nodeWithVal != nil{
		node.nextNode = nodeWithVal.nextNode
		nodeWithVal.nextNode = node
	}

}


// Doubly linked list -------------------


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



func (linkedList *LinkedList) AddToHeadDoubly(property int){
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil{
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}


	linkedList.headNode = node
}

func (linkedList *LinkedList) AddAfterDoubly(nodeProp, property int){
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWithVal *Node
	nodeWithVal = linkedList.NodeWithVal(nodeProp)

	if nodeWithVal != nil{
		node.nextNode = nodeWithVal.nextNode
		fmt.Println("node With nextNode ---->",node.nextNode)

		node.previousNode = nodeWithVal
		fmt.Println("node With previousNode ---->",node.previousNode)

		nodeWithVal.nextNode = node
		fmt.Println("nodeWithVal next node ---->",nodeWithVal.nextNode)
	}
}

func(linkedList *LinkedList) AddToEndDoubly(property int){
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node 
	lastNode = linkedList.LastNode()

	if lastNode != nil{
		lastNode.nextNode = node 
		node.previousNode = lastNode
	}
}


func main(){
	var ll LinkedList

	ll = LinkedList{}
	ll.AddToHeadDoubly(1)
	ll.AddToHeadDoubly(3)
	ll.AddToHeadDoubly(5)

	// ll.AddToHead(5)
	// ll.AddToHead(7)

	ll.AddAfterDoubly(1,15)

	ll.AddToEndDoubly(9)
	
	// ll.AddToEnd(5)
	// ll.AddAfter(1,15)
	ll.IterateList()

	

	fmt.Println(ll.headNode.property)

	var node *Node
	node = ll.NodeBetweenValues(1,3)
	fmt.Println(node.property)

	


}