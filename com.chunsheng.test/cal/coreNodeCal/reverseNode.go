package main

import (
	"com.chunsheng.test/cal/coreStruct"
)

func Way1(node *coreStruct.Node) *coreStruct.Node {
	current := node
	var mid *coreStruct.Node
	for{
		midNow := mid
		mid = current
		current = current.Next
		mid.Next = midNow
		if current == nil{
			break
		}
	}
	return mid
}


func Way2(node *coreStruct.Node) *coreStruct.Node{
	if node == nil{
		return nil
	}
	stack := &coreStruct.Stack{}
	// 将元素入栈
	current := node
	for{
		stack.Push(current)
		if current.Next == nil{
			break
		}
		current = current.Next
	}
	// 将元素出栈
	resp := stack.Pop().(*coreStruct.Node)
	curr := resp
	for {
		if stack.CLen == 0{
			break
		}
		nxt := stack.Pop().(*coreStruct.Node)
		nxt.Next = nil
		curr.Next = nxt
		curr = nxt
	}
	return resp
}



func main() {
	myNode := coreStruct.GetANodesFromArray([]int{1,2,3,4,5,6,7,8})
	coreStruct.PrintNode(myNode)
	mn := Way2(myNode)
	coreStruct.PrintNode(mn)
}

