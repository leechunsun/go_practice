package coreStruct

import "sync"

type interNode struct {
	prefix *interNode
	val interface{}
	next *interNode
}

type Stack struct {
	CLen int64  // 当前栈内元素的长度
	Top interface{}  // 栈顶元素
	Low interface{}  // 栈底元素
	iNode *interNode  // 栈列表
	tailNode *interNode  // 栈顶指针
	mu sync.Mutex  // 线程锁
	ThreadSafe bool  // 是否线程安全
}

func (s *Stack) Push (it interface{}){
	if s.ThreadSafe{
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	var newNode *interNode
	if s.iNode == nil{
		newNode = &interNode{prefix:nil, val:it, next:nil}
		s.iNode = newNode
		s.Low = it

	}else {
		newNode = &interNode{prefix:s.tailNode, val:it, next:nil}
		s.tailNode.next = newNode
	}
	s.Top = it
	s.tailNode = newNode
	s.CLen += 1
}

func (s *Stack) Pop() interface{} {
	if s.ThreadSafe{
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	if s.CLen <= 0{
		panic("stack empty...")
	}
	val := s.tailNode.val
	prefix := s.tailNode.prefix
	if prefix == nil{
		s.Top = nil
		s.Low = nil
		s.tailNode = nil
	} else {
		prefix.next = nil
		s.tailNode = prefix
		s.Top = prefix.val
	}
	s.CLen -= 1
	return val
}
