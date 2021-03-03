package coreStruct

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val  int
	Next *Node
}


func GetANodesFromArray(arr []int) *Node {
	if len(arr) <= 0 {
		return nil
	}
	Base := &Node{Val:arr[0]}
	CurrentNode := Base
	for _, val := range arr[1:] {
		this := &Node{Val:val}
		CurrentNode.Next = this
		CurrentNode = this
	}
	return Base
}

func PrintNode(node *Node){
	if node == nil{
		fmt.Println("current node is nil....")
		return
	}
	Curr := node
	mystr := ""
	for{
		cnext := Curr.Next
		mystr += strconv.Itoa(Curr.Val)
		if cnext == nil{
			break
		}
		mystr += " -> "
		Curr = cnext
	}
	fmt.Println(mystr)
}




