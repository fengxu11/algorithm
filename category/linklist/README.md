### 链表 LinkList

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
