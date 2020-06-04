package list

import (
	"fmt"
)
type Node struct {
	Data string
	Link *Node
}
type LinkedList struct {
	First *Node
	Last  *Node
}
func (list *LinkedList) Add(data string) {
	var node Node
	node.Data = data
	
	if list.First == nil {
		node.Link = nil
		list.First = &node
	} else {
		list.Last.Link = &node
	}
	list.Last = &node
}

func (list LinkedList) PrintList() {

	temporal := list.First
	for temporal != nil {
		fmt.Println(temporal.Data)
		temporal = temporal.Link
	}
}
