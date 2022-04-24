package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	root := Ints2TreeNode([]int{7, 3, 15, null, null, 9, 20})
	obj := Constructor(root)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param2 := obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param2 = obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
}
