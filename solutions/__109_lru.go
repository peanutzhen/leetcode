package main

import "fmt"

// Node 双向链表节点定义
type Node struct {
	key        int // 用于反向获取Cache指针
	value      int
	prev, next *Node
}

type LRUCache struct {
	hashmap  map[int]*Node
	length   int
	capacity int
	head     *Node // 最近最常用key-value指针
	tail     *Node // 最近最不常用key-value指针
}

func Constructor(capacity int) LRUCache {
	rtv := LRUCache{
		hashmap:  make(map[int]*Node),
		capacity: capacity,
		head:     &Node{}, // 虚拟头节点
		tail:     &Node{}, // 虚拟尾节点 便于insert/delete操作
	}
	rtv.head.next = rtv.tail
	rtv.tail.prev = rtv.head
	return rtv
}

func (this *LRUCache) removeNode(node *Node) {
	p, n := node.prev, node.next
	p.next = n
	n.prev = p
	node.next, node.prev = nil, nil
	this.hashmap[node.key] = nil
	this.length--
}

func (this *LRUCache) addNodeToHead(node *Node) {
	p := this.head.next
	this.head.next = node
	node.next = p
	node.prev = this.head
	p.prev = node
	this.length++
}

func (this *LRUCache) Get(key int) int {
	node := this.hashmap[key]
	if node == nil {
		return -1
	}
	// 放到链表头部 表示最近使用过
	this.removeNode(node)
	this.addNodeToHead(node)
	this.hashmap[key] = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.hashmap[key] != nil {
		this.hashmap[key].value = value
		this.Get(key) // Tips: 利用Get放到队头
		return
	}
	if this.length == this.capacity {
		this.removeNode(this.tail.prev) // 先移除最近最少使用的
	}
	// 添加节点
	node := &Node{
		key:   key,
		value: value,
	}
	this.addNodeToHead(node)
	this.hashmap[key] = node
}

func main() {
	//["LRUCache","put","get","put","get","get"]
	//[[1],[2,1],[2],[3,2],[2],[3]]
	obj := Constructor(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Put(3, 2)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
}
