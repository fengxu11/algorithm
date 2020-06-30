### 链表 LinkList

> single.go

```go

// 数据存储在节点中
// 最后一根节点的Next 指向的是空
type Node struct {
    Data interface{}
    Next Node
}

```

> 数据存储在节点中

> 最后一根节点的Next 指向的是空

> 链表是一个真正的动态的数据结构、不需要处理固定容量的问题


```

添加操作
addLast(data)       O(n)
addFirst(data)      O(1)
add(index, data)    O(n/2) = O(n)


删除
removeLast(data)    O(n)
removeFirst(data)   O(n)
remove(index, data)    O(n/2) = O(n)


查找
get(index)          O(n)
contains(data)      O(n)


对于链表来说、时间复杂度
    增: O(n)
    删: O(n)
    改: O(n)
    查: O(n)

```


### 使用链表 实现栈 

> stack.go



### 使用链表、带尾指针的链表 实现队列

> queue.go

```

    1. 从链表头 删除元素非常容易、 所以head作为队列的 队首
    2. 从链表尾 删除元素不容易、所以tail 作为队列的 队尾
    3. 从 head删除元素、 从tail插入元素

```


### LeetCode 203. remove-linked-list-elements
```

删除链表中 等于给定值 val的所有节点

示例: 
    输入: 1->2->6->3->4->5->6, val = 6
    输出: 1->2->3->4->5

代码在: 203_remove-element_test.go

```