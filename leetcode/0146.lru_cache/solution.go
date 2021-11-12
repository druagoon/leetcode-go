// 146. LRU 缓存机制(https://leetcode-cn.com/problems/lru-cache/)

package leetcode

// leetcode submit region begin(Prohibit modification and deletion)
type DoublyListNode struct {
	key   int
	value int
	prev  *DoublyListNode
	next  *DoublyListNode
}

func initDoublyListNode(key, value int) *DoublyListNode {
	return &DoublyListNode{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DoublyListNode
	head     *DoublyListNode
	tail     *DoublyListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    map[int]*DoublyListNode{},
		head:     initDoublyListNode(0, 0),
		tail:     initDoublyListNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {

	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	node := initDoublyListNode(key, value)
	this.cache[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
		this.size--
	}
}

func (this *LRUCache) addToHead(node *DoublyListNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoublyListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DoublyListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DoublyListNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// leetcode submit region end(Prohibit modification and deletion)
