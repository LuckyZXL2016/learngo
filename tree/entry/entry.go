package main

import (
	"fmt"
	"learngo/tree"
)

// 扩充类型方法1：组合方式
type myTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}

	fmt.Println(nodes)
	var root tree.Node
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	// 使用new与上面的效果一致
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	// 值接收者传地址也可以，自动取值出来使用
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()
	// nil指针也可以调用方法
	fmt.Println("use nil ptr:")
	var pRoot2 *tree.Node
	pRoot2.SetValue(200)
	pRoot2 = &root
	pRoot2.SetValue(300)
	pRoot2.Print()

	fmt.Println("Pre Traversing:")
	root.Traverse()
	fmt.Println()

	// 扩充类型方法1：组合方式
	fmt.Println("Post Traversing:")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
