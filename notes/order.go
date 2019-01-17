package main

import "fmt"

func main() {
	//二叉树遍历
	arr := []int{10, 5, 24, 33, 60, 40, 45, 15, 27, 49, 23, 42, 56, 12, 8, 55, 2, 9}
	fmt.Println(arr)
	t := creatTree(arr)
	preorder(*t)
	fmt.Println(Preorder)
	inorder(*t)
	fmt.Println(Inorder)
	postorder(*t)
	fmt.Println(Postorder)

	levelorder(*t)
	fmt.Println(Levelorder)
}

type tree struct {
	data  int
	left  *tree
	right *tree
}

var Preorder, Inorder, Postorder = []int{}, []int{}, []int{}

//´´½¨¶þ²æÊ÷
func creatTree(arr []int) *tree {
	// d := make([]tree, 0)
	// for i, ar := range arr {
	// 	d = append(d, tree{})
	// 	d[i].data = ar
	// }
	// for i := 0; i < len(arr)/2; i++ {
	// 	d[i].left = &d[i*2+1]
	// 	if i*2+2 < len(d) {
	// 		d[i].right = &d[i*2+2]
	// 	}
	// }
	// fmt.Println(d)
	// return d

	var root *tree
	for _, v := range arr {
		root = add(root, v)
	}
	return root
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.data = value
		return t
	}
	if value < t.data {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//Ç°Ðò±éÀú
func preorder(root tree) {
	Preorder = append(Preorder, root.data)
	if root.left != nil {
		preorder(*root.left)
	}
	if root.right != nil {
		preorder(*root.right)
	}
}

//ÖÐÐò±éÀú
func inorder(root tree) {
	if root.left != nil {
		inorder(*root.left)
	}
	Inorder = append(Inorder, root.data)
	if root.right != nil {
		inorder(*root.right)
	}
}

//ºóÐò±éÀú
func postorder(root tree) {
	if root.left != nil {
		postorder(*root.left)
	}
	if root.right != nil {
		postorder(*root.right)
	}
	Postorder = append(Postorder, root.data)
}

//¹ã¶ÈÓÅÏÈ±éÀú
var Levelorder = [][]int{}

var Level []*tree
var i, s, j, k int = 0, 0, 0, 1

func levelorder(root tree) [][]int {
	Level = append(Level, &root)
	for i = 0; ; i++ {
		j = s
		s = s + k
		k = 0
		var Levelorder0 = []int{}
		for ; j < s; j++ {
			Levelorder0 = append(Levelorder0, Level[j].data)
			if Level[j].left != nil {
				k++
				Level = append(Level, Level[j].left)
			}
			if Level[j].right != nil {
				k++
				Level = append(Level, Level[j].right)
			}
		}
		Levelorder = append(Levelorder, Levelorder0)
		if k == 0 {
			break
		}
	}
	return Levelorder
}
