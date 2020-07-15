
### Queue

```

    1. 队列也是一种线性结构
    2. 相比数组, 队列对应的操作是数组的子集
    3. 只能从 一端(队尾) 添加元素, 只能从另一端(队首)取出元素

    *队列是一种先进先出的数据结构
    *First In First Out (FIFO)

```


### 数组队列
> ArrayQueue 数组队列
> 	GetSize() int                   时间复杂度是 O(1)
>	IsEmpty() bool                  时间复杂度是 O(1)
>	Enqueue(e interface{})          时间复杂度是 O(1) 均摊情况下
>	Dequeue() interface{}           时间复杂度是 O(n) 数据量越大、越慢
>	GetFront() interface{}          时间复杂度是 O(1)

### 循环队列
```

基于 数组实现
    1. 要有两个指向 头和尾、也就是 front tail
    2. front == tail 队列则为空
    3. (tail + 1) % capacity  == front 则队列满了
    4. 这么做、浪费了1个空间

```

> ArrayQueue 数组队列
> 	GetSize() int                   时间复杂度是 O(1)
>	IsEmpty() bool                  时间复杂度是 O(1)
>	Enqueue(e interface{})          时间复杂度是 O(1) 均摊情况下
>	Dequeue() interface{}           时间复杂度是 O(1) 均摊情况下
>	GetFront() interface{}          时间复杂度是 O(1)
