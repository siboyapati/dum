package Design

const (
	mod = 1024
)

type listNode struct {
	key   int
	value int
	next  *listNode
}
type MyHashMap struct {
	set [mod]*listNode
}

func Constructor() MyHashMap {
	arr := [mod]*listNode{}
	return MyHashMap{arr}
}

func (this *MyHashMap) Put(key int, value int) {
	i := key % mod
	ptr := this.set[i]
	for ptr != nil {
		if ptr.key == key {
			ptr.value = value
			return
		} else {
			ptr = ptr.next
		}
	}
	newNode := &listNode{key: key, value: value, next: ptr}
	this.set[i] = newNode
}

func (this *MyHashMap) Get(key int) int {
	i := key % mod
	ptr := this.set[i]
	for ptr != nil {
		if ptr.key == key {
			return ptr.value
		} else {
			ptr = ptr.next
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	i := key % mod
	ptr := this.set[i]
	prev := &listNode{next: ptr}
	head := prev
	for ptr != nil {
		if ptr.key == key {
			prev.next = ptr.next
			break
		} else {
			ptr = ptr.next
			prev = prev.next
		}
	}
	this.set[i] = head.next
}
