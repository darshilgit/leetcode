package linkedlist

type LRUCache struct {
	lnode    *Node
	rnode    *Node
	cap      int
	size     int
	key2node map[int]*Node
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {

	lnode := &Node{}
	rnode := &Node{}

	lnode.next = rnode
	lnode.prev = nil

	rnode.next = nil
	rnode.prev = lnode

	return LRUCache{
		lnode:    lnode,
		rnode:    rnode,
		cap:      capacity,
		size:     0,
		key2node: make(map[int]*Node, capacity),
	}

}

func (this *LRUCache) Get(key int) int {

	if nodeptr, ok := this.key2node[key]; ok {

		// remove the nodeptr inplace
		nodeptr.prev.next = nodeptr.next
		nodeptr.next.prev = nodeptr.prev
		nodeptr.next = nil
		nodeptr.prev = nil

		// shift node to the right
		this.rnode.prev.next = nodeptr
		nodeptr.next = this.rnode
		nodeptr.prev = this.rnode.prev
		this.rnode.prev = nodeptr

		return nodeptr.val

	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	nodeval := this.Get(key)
	if nodeval == -1 {
		if this.size < this.cap {
			//only add
			dummy := &Node{
				key:  key,
				val:  value,
				prev: nil,
				next: nil,
			}
			// update next and prev ptrs for dummy
			this.rnode.prev.next = dummy
			dummy.next = this.rnode
			dummy.prev = this.rnode.prev
			this.rnode.prev = dummy

			//add to map
			this.key2node[key] = dummy
			this.size++

		} else {
			//evict and add
			nodeptr := this.lnode.next
			nodeptr.prev.next = nodeptr.next
			nodeptr.next.prev = nodeptr.prev

			nodeptr.next = nil
			nodeptr.prev = nil

			//delete from map
			delete(this.key2node, nodeptr.key)

			//add to the right
			nodeptr.key = key
			nodeptr.val = value
			dummy := nodeptr

			this.rnode.prev.next = dummy
			dummy.next = this.rnode
			dummy.prev = this.rnode.prev
			this.rnode.prev = dummy

			//add to map
			this.key2node[key] = dummy

		}
	} else {
		// inplace update for value
		this.rnode.prev.val = value
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
