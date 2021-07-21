package main

import "fmt"

type Node struct {
	key string
	left *Node
	right *Node
}

func PreOrder(root *Node){

	fmt.Println(root.key)

	if root.left != nil {
		PreOrder(root.left)
	}

	if root.right != nil {
		PreOrder(root.right)
	}
}

func PostOrder(root *Node){

	if root.left != nil {
		PostOrder(root.left)
	}

	if root.right != nil {
		PostOrder(root.right)
	}
	fmt.Println(root.key)
}

func main(){

	var inputExpression string

	fmt.Scan(&inputExpression)

	for i:=0; i<len(inputExpression); i++{
		if inputExpression == "+" || inputExpression == "-" || inputExpression == "*" || inputExpression == "/"{
			fmt.Println("We have encountered an operator: Push into the stack")

		}else {
			fmt.Println("We have encountered an operand, POP the last operator from the stack")
		}
	}

}