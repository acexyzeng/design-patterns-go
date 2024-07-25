package lfu

import "container/list"

// least frequently used
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (l *LFUCache) Get(key int) int {
	value, ok := l.nodes[key]
	if !ok {
		return -1
	}

	currentNode := value.Value.(*node)
	l.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := l.lists[currentNode.frequency]; !ok {
		l.lists[currentNode.frequency] = list.New()
	}

	newList := l.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	l.nodes[key] = newNode
	if currentNode.frequency-1 == l.min && l.lists[currentNode.frequency-1].Len() == 0 {
		l.min++
	}
	return currentNode.value
}

func (l *LFUCache) Put(key, value int) {
	if l.capacity == 0 {
		return
	}

	if currentValue, ok := l.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		l.Get(key)
		return
	}

	if l.capacity == len(l.nodes) {
		currentList := l.lists[l.min]
		backNode := currentList.Back()
		delete(l.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}

	l.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}

	if _, ok := l.lists[1]; !ok {
		l.lists[1] = list.New()
	}
	newList := l.lists[1]
	newNode := newList.PushFront(currentNode)
	l.nodes[key] = newNode
}
