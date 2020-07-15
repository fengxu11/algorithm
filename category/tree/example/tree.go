package example

import (
	"bytes"
	"fmt"
	queue "github.com/algorithm/category/linklist/example"
	stack "github.com/algorithm/category/stack/example"

	"github.com/algorithm/category/common"
)

// Node 节点
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

// BST 二分搜索树
type BST struct {
	root *Node
	size int
}

// NewBST  创建二分搜索树
func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

// GetSize get size
func (b *BST) GetSize() int {
	return b.size
}

// IsEmpty 判断二分搜索树 是否为空
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// Add 添加元素
func (b *BST) Add(e interface{}) {

	b.root = b.add(b.root, e)
}

// 往 以node为根的二分搜索树中、插入元素e、递归实现
func (b *BST) add(node *Node, e interface{}) *Node {

	// 递归终止条件
	if node == nil {
		b.size++
		return &Node{
			Data: e,
		}
	}

	// 要插入的元素 和 树节点中的元素进行比较、小于则插入到左子树、大于插入到右子树
	if common.Compare(e, node.Data) < 0 {
		node.Left = b.add(node.Left, e)
	} else if common.Compare(e, node.Data) > 0 {
		node.Right = b.add(node.Right, e)
	}

	return node
}

// Contains 查看二分搜索树中 是否包含元素e
func (b *BST) Contains(e interface{}) bool {

	return b.contains(b.root, e)
}

// 递归、以node为根的二分搜索树中 是否包含元素e
func (b *BST) contains(node *Node, e interface{}) bool {

	if node == nil {
		return false
	}

	if common.Compare(e, node.Data) == 0 {
		return true
	} else if common.Compare(e, node.Data) < 0 {
		return b.contains(node.Left, e)
	} else {
		return b.contains(node.Right, e)
	}
}

// PreOrder 前序遍历
func (b *BST) PreOrder() {
	preOrder(b.root)
}

func preOrder(node *Node) {

	if node != nil {
		fmt.Println(node.Data)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

// InOrder 中序遍历
func (b *BST) InOrder() {
	inOrder(b.root)
}

func inOrder(node *Node) {

	if node != nil {
		preOrder(node.Left)
		fmt.Println(node.Data)
		preOrder(node.Right)
	}
}

// PostOrder 后序遍历
func (b *BST) PostOrder() {
	postOrder(b.root)
}

func postOrder(node *Node) {

	if node != nil {
		preOrder(node.Left)
		preOrder(node.Right)
		fmt.Println(node.Data)
	}
}

func (b *BST) String() string {
	var buffer bytes.Buffer
	// 前序打印
	//generatePreOrderBSTString(b.root, 0, &buffer)

	// 中序打印
	//generateInOrderBSTString(b.root, 0, &buffer)

	// 后续打印
	generatePostOrderBSTString(b.root, 0, &buffer)

	return buffer.String()
}

func generatePostOrderBSTString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "null \n ")
		return
	}

	generatePostOrderBSTString(node.Left, depth+1, buffer)
	generatePostOrderBSTString(node.Right, depth+1, buffer)
	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(node.Data) + "\n")
}

func generateInOrderBSTString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "null \n ")
		return
	}

	generateInOrderBSTString(node.Left, depth+1, buffer)
	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(node.Data) + "\n")
	generateInOrderBSTString(node.Right, depth+1, buffer)
}

func generatePreOrderBSTString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "null \n ")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(node.Data) + "\n")
	generatePreOrderBSTString(node.Left, depth+1, buffer)
	generatePreOrderBSTString(node.Right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

// 非递归 二分搜索树 前序遍历、使用栈模拟系统栈 进行非递归遍历
func (b *BST) preOrderNR() {
	stack := stack.New()
	stack.Push(b.root)
	// 栈不为空、就一直循环
	for !stack.IsEmpty() {
		cur := stack.Pop()
		if node, ok := cur.(*Node); ok {
			// 打印
			fmt.Println(node.Data)
			// 访问左子树
			if node.Right != nil {
				stack.Push(node.Right)
			}
			if node.Left != nil {
				stack.Push(node.Left)
			}
		}
	}
}

// 层序遍历 一层一层的遍历、从根节点遍历、根节点的层级一般定义为 0
// 使用 链表进行 层序遍历
func (b *BST) levelOrder() {

	q := queue.NewLinkListQueue()
	q.Enqueue(b.root)

	for !q.IsEmpty() {
		cur := q.Dequeue()
		if node, ok := cur.(*Node); ok {
			fmt.Println(node.Data)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}
}

