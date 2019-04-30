class DLLNode:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class DLList:
    def __init__(self):
        self.head = DLLNode(0, 0)  # Dummy Head
        self.tail = DLLNode(0, 0)  # Dummy Tail
        self.head.next = self.tail
        self.tail.prev = self.head

    def insert_after(self, prev_node, new_node):
        # prev_node is node to insert after, new_node is node to be inserted after

        # 1. Check if the given prev_node is None
        if prev_node == None:
            print("Previous node cannot be null!")
            return
        # 2. Make next of new node as next of prev node
        new_node.next = prev_node.next
        # 3. Make prev_node as previous of new_node
        prev_node.next = new_node
        # 4. Make prev_node as previous of new_node
        new_node.prev = prev_node
        # 5. Change previous of new_nodes' next node
        if new_node.next is not None:
            new_node.next.prev = new_node

    def remove(self, node):
        prev = node.prev
        next = node.next
        prev.next = next
        next.prev = prev
        node.prev, node.next = None, None

    def remove_from_head(self):
        delete = self.head.next  # delete: node to be deleted
        self.remove(delete)
        return delete.key

    def move_to_end(self, node):
        if self.head == node:
            self.remove_from_head()
        else:
            self.remove(node)
        self.insert_after(self.tail.prev, node)

    def add_to_end(self, node):
        self.insert_after(self.tail.prev, node)


class LRUCache:

    def __init__(self, capacity):
        self.capacity = capacity
        self.dll = DLList()
        self.hashmap = {}

    def get(self, key):
        if key in self.hashmap:
            cache_node = self.hashmap[key]
            self.dll.move_to_end(cache_node)
            return cache_node.value
        return -1

    def put(self, key, value):
        if key in self.hashmap:
            cache_node = self.hashmap[key]
            cache_node.value = value
            self.dll.move_to_end(cache_node)
        else:
            cache_node = DLLNode(key, value)
            self.dll.add_to_end(cache_node)
            self.hashmap[key] = cache_node
        if len(self.hashmap) > self.capacity:
            delete_key = self.dll.remove_from_head()
            del self.hashmap[delete_key]

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)