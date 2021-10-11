package main

import "fmt"

//node class with property int and nextNode which points to the Node class
type Node struct{
	property int
	nextNode *Node
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

func main(){
	var ll LinkedList

	ll = LinkedList{}
	ll.AddToHead(1)
	ll.AddToHead(3)

	fmt.Println(ll.headNode.property)

}