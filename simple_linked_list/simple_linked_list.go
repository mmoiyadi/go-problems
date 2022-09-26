package linkedlist

import "errors"

// Define the List and Element types here.
type Node struct{
	Next *Node
	Value int
}
type List struct{
    Head *Node
		Length int
}

func New(nums []int) *List {
	
	var l = &List{Head: nil, Length: 0}
	for _, elem := range nums{
        l.Push(elem)
    } 
	return l
}

func (l *List) Size() int {
	if l == nil{
		return 0
	}
	return l.Length
}

func (l *List) Push(element int) {
	node := &Node{Value: element, Next: nil}
	head := l.Head
	node.Next = head
	l.Head = node
	l.Length += 1
}

func (l *List) Pop() (int, error) {
	if l.Head == nil{
		return 0, errors.New("Nothing to pop")
	}
	poppedItem := l.Head
	l.Head = poppedItem.Next
	l.Length -= 1
	return poppedItem.Value, nil
}

func (l *List) Array() []int {
	arr := make([]int, 0)
	node := l.Head
	for node != nil{
		arr = append([]int{node.Value}, arr... )
		node = node.Next
	}
	return arr
}

func (l *List) Reverse() *List {
	reversed := New([]int{})
	node := l.Head
	for node != nil{
		reversed.Push(node.Value)
		node = node.Next
	}
	return reversed
}
