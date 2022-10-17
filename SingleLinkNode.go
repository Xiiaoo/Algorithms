package main

import (
	"fmt"
)

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func ShowNode(n *LinkNode) {
	for n != nil {
		fmt.Println(n.Data)
		n = n.NextNode
	}
}

func DeleteNode(n *LinkNode){
	
}

func main() {
	node := new(LinkNode)
	node.Data = 1

	var temp *LinkNode
	temp=node

	node1 := new(LinkNode)
	node1.Data = 2
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 3
	node1.NextNode = node2

	// 头插
	// for i := 0; i < 5; i++ {
	// 	newnode := LinkNode{Data: int64(i)}
	// 	newnode.NextNode = node
	// 	node = &newnode
	// }

	// 尾插
	var last *LinkNode
	last=node2
	for i := 0; i < 5; i++ {
		newnode := LinkNode{Data: int64(i)}
		(*last).NextNode=&newnode
		last=&newnode
	}

	fmt.Println(temp)
	fmt.Println(temp.NextNode)

	// ShowNode(node)
}
