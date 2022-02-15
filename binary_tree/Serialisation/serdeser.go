package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	str := ser(root)
	return str[:len(str)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	return deser(&str)
}

func main() {

	l5 := TreeNode{Val: 5, Left: nil, Right: nil}
	l4 := TreeNode{Val: 4, Left: nil, Right: nil}
	l3 := TreeNode{Val: 3, Left: &l4, Right: &l5}
	l2 := TreeNode{Val: 2, Left: nil, Right: nil}
	l1 := TreeNode{Val: 1, Left: &l2, Right: &l3}

	codec := Constructor()
	//var root = codec.deserialize("1,2,3,null,null,4,5")
	var string = codec.serialize(&l1)

	fmt.Println(string)
	fmt.Println(codec.deserialize(string))
}

func ser(root *TreeNode) string {
	if root == nil {
		return "null,"
	}
	return fmt.Sprintf("%d,", root.Val) + ser(root.Left) + ser(root.Right)
}

func deser(str *[]string) *TreeNode {
	if len(*str) == 0 {
		return nil
	}

	v, err := strconv.ParseInt((*str)[0], 10, 32)

	*str = (*str)[1:]

	if err == nil {

		left := deser(str)
		right := deser(str)

		return &TreeNode{
			Val:   int(v),
			Left:  left,
			Right: right,
		}
	}

	return nil
}
