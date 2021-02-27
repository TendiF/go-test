package BinaryTree

import (
	"testing"
)


func TestSearch(t *testing.T){
	root := Node{Key: 10}

	root.Insert(20)
	root.Insert(1)
	root.Insert(2)
	root.Insert(40)
	root.Insert(30)
	root.Insert(60)

	t.Log(root.Search(60))

	if !root.Search(60) {
		t.Error("60 should true")
	}

	if root.Search(100) {
		t.Error("100 should false")
	}
}