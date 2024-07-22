package LRU

// least recently used cache
type LRUCache struct {
	Cap        int
	Keys       map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Cap:  cap,
		Keys: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.Keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node := &Node{key: key, val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}

	if node == this.tail {
		this.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}
