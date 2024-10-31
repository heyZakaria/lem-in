package stack

import(
	"fmt"
)

type node struct{
	val string
	next *node
}

func (n *node) GetValue() string{
	return n.val
}

type Stack struct{
	Head *node
	Tail *node
}

func NewStack() *Stack {
	return &Stack{
		Head: nil,
		Tail: nil,
	}
}

func (s *Stack) Push(val_ string){
	node := &node{
		val: val_,
		next: nil,
	}
	if s.Head == nil {
		s.Head = node
		s.Tail = s.Head
		s.Head.next = s.Tail
		s.Tail.next = nil
		return
	}
	s.Tail.next = node
	s.Tail = node
}

func (s *Stack) Show(){
	current := s.Head
	for current != nil{
		fmt.Printf("-> %v ",current.val)
		current = current.next
	}
	fmt.Println()
}

func (s *Stack) ConvertToArray() []string {
	array := make([]string,0)
	current := s.Head
	for current != nil{
		array = append(array, current.val)
		current = current.next
	}
	return array
}

func (s *Stack) Pop() *node{
	if s.IsNull() {
		return nil
	}
	node := s.Head
	if s.Head == s.Tail {
		s.Head.next = nil
		s.Head = nil
		s.Tail = nil
		return node
	}
	current := s.Head
	for current != nil {
		if current.next == s.Tail {
			node = s.Tail
			current.next = nil
			s.Tail = current
		}
		current = current.next
	}
	return node
}

func (s *Stack) IsNull() bool{
	return (s.Head == nil)
}
