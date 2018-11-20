package main

import "fmt"

type treeNode struct{
	value int
	left,right *treeNode
}

//Go 语言，方法，定义在结构体外,(node treeNode) 称为接收者,
//注意，值传递
func (node treeNode) print(){
	fmt.Println(node.value)
}
//要定义修改方法，但是调用的额时候，都可以使用值调用
func (node *treeNode) setValue(value int){
	node.value = value
}


func createNode(value int) *treeNode{
	return &treeNode{value:value}
}
func main() {
	var root treeNode
	fmt.Println(root)

	root1 := treeNode{value:3}
	root1.left = &treeNode{}
	root1.left.right = createNode(2)
	root1.right = &treeNode{5,nil,nil}
	fmt.Println(root1)
	root1.right.left = createNode(0)
	nodes := []treeNode{
		{value:3},
		{},
	}
	fmt.Println(nodes)

	//调用结构体的方法
	root1.print()
	root1.setValue(4)
	root1.print()

	//nil 调用方法
}
