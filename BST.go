package main

import (
	"errors"
	"fmt"
)

//Node represent components of BST
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert mhetod to add a node to the tree
func (n *Node) Insert(k int) error {
	//if tree is empty - return eror
	if n == nil {
		return errors.New("[Error] tree is nil")
	}

	//if inserted value already exists return corresponding error
	if n.Key == k {
		return errors.New("[Error] value already exists")
	}

	//move right when insert value larger than the Key, if node empty make the new node with the Key - K, otherwise call this mhetod agian
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
			return nil
		} else {
			return n.Right.Insert(k)
		}
		//same as before
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
			return nil
		} else {
			return n.Left.Insert(k)
		}
	}
	return nil
}

//Search method
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)

	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

//Traversing the tree in sort order
func (n *Node) Traverse(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.Traverse(f)
	f(n)
	n.Right.Traverse(f)
}

//Subfunction for finding min value of the BST
func minValue(n *Node) *Node {
	temp := n
	for nil != temp && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

//Delete function
func Delete(n *Node, k int) *Node {
	if n == nil {
		return n
	}

	if n.Key > k {
		n.Left = Delete(n.Left, k)
	}
	if n.Key < k {
		n.Right = Delete(n.Right, k)
	}
	if n.Key == k {
		if n.Left == nil && n.Right == nil {
			n = nil
			return n
		}
		if n.Left == nil && n.Right != nil {
			temp := n.Right
			n = nil
			n = temp
			return n
		}
		if n.Left != nil && n.Right == nil {
			temp := n.Left
			n = nil
			n = temp
			return n
		}
		tempNode := minValue(n.Right)
		n.Key = tempNode.Key
		n.Right = Delete(n.Right, tempNode.Key)
	}
	return n
}

func main() {
	tree := Node{Key: 10}
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(101)
	tree.Insert(105)
	tree.Insert(108)
	tree.Insert(2)
	tree.Insert(20)
	tree.Traverse(func(n *Node) { fmt.Printf("%v\n", n.Key) })
	fmt.Println("=============================================")
	Delete(&tree, 20)
	Delete(&tree, 108)
	Delete(&tree, 2)
	Delete(&tree, 10)
	tree.Traverse(func(n *Node) { fmt.Printf("%v\n", n.Key) })

}

//insert method works recursively:
//Compare key with the reciver and go to the cihld on the right if its larger,
//or go to the left if its smaller
//If the child is empty, put new node there
//If not empty, call insert(k) again

//search also recursive:
//this method returns either true when value of the node is found
//and false if the node is not found
