package leetcode

// DLinkedNode ...
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// LRUCache ...
type LRUCache struct {
	capacity, size int
	head, tail     *DLinkedNode
	entry          map[int]*DLinkedNode
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		entry:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get ...
func (c *LRUCache) Get(key int) int {
	if node, ok := c.entry[key]; ok {
		c.moveToHead(node)
		return node.value
	}

	return -1
}

// Put ...
func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.entry[key]; ok {
		node := c.entry[key]
		node.value = value
		c.moveToHead(node)
		return
	}

	node := &DLinkedNode{
		key:   key,
		value: value,
	}
	c.entry[key] = node
	c.addToHead(node)
	c.size++
	if c.size > c.capacity {
		removed := c.removeTail()
		delete(c.entry, removed.key)
		c.size--
	}
}

func (c *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToHead(node *DLinkedNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *DLinkedNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
