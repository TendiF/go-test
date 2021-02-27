package BinaryTree

import (
	"fmt"
)

type Node struct {
	Key int
	Left *Node
	Right *Node
}

//Insert
func (n *Node) Insert (k int) {
	if n.Key < k {
		//go to right
		if n.Right == nil {
			n.Right = &Node{Key : k}
		}else{
			n.Right.Insert(k)
		}
	} else if n.Key >= k {
		// go to left
		if n.Left == nil {
			n.Left = &Node{Key : k}
		}else{
			n.Left.Insert(k)
		}
	} 
}

//Search
func (n *Node) Search (k int) bool{
	if n == nil {
		return false
	}
	  
	if n.Key < k {
		//go right
		return n.Right.Search(k)
	} else if n.Key > k {
		//go left
		return n.Left.Search(k)
 	}

	return true
}

func (n *Node)PreOrder(){
	// val, left , right
	if n != nil {
		fmt.Println("PreOrder", n.Key)
		if n.Left != nil {
			n.Left.PreOrder()
		}
		if n.Right != nil {
			n.Right.PreOrder()
		}
	}
}

func (n *Node)InOrder(){
	// left, val , right
	if n != nil {
		if n.Left != nil {
			n.Left.InOrder()
		}
		fmt.Println("InOrder", n.Key)
		if n.Right != nil {
			n.Right.InOrder()
		}
	}
}

func (n *Node)PostOrder(){
	// left, right , val
	if n != nil {
		if n.Left != nil {
			n.Left.PostOrder()
		}
		if n.Right != nil {
			n.Right.PostOrder()
		}
		fmt.Println("PostOrder", n.Key)
	}
}

func RunBinary (){
	twenty := Node{Key: 20, Left: &Node{Key:40}, Right: &Node{Key: 60} }
	thirty := Node{Key: 30, Left: &Node{Key:50}}
	root := Node{Key: 10, Left: &twenty, Right: &thirty}

	root.PreOrder()
	root.PostOrder()
	root.InOrder()
}