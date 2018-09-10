package tree

import "fmt"

// 定义结构体
type Node struct {
	Value       int
	Left, Right *Node
}

// 打印节点的值，(node node)是接收者，表示Print()方法是给node来用的
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 设置节点的值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

// 自定义工厂函数：相当于自己写的构造函数。go没有构造函数
func CreateNode(value int) *Node {
	// 可以返回局部变量的地址
	return &Node{Value: value}
}
