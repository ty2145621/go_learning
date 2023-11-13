package real

type LFUCache struct {
	cache                    map[int]*Node
	freq                     map[int]*DLinkList
	numMax, numUsed, minFreq int
}

func New(k int) *LFUCache {
	return &LFUCache{
		cache:  map[int]*Node{},
		freq:   map[int]*DLinkList{},
		numMax: k,
	}
}

func (l *LFUCache) get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.incrFreq(node)
		return node.val
	}
	return -1
}

func (l *LFUCache) set(key, val int) {
	if node, ok := l.cache[key]; ok {
		node.val = val
		l.incrFreq(node)
		return
	}
	node := &Node{
		key:  key,
		val:  val,
		freq: 1,
	}
	if l.numUsed == l.numMax {
		key := l.freq[l.minFreq].removeTail()
		delete(l.cache, key)
		l.numUsed--
	}
	l.minFreq = 1
	l.numUsed++
	l.cache[key] = node
	if _, ok := l.freq[1]; !ok {
		l.freq[1] = NewDLinkList()
	}
	l.freq[1].pushFront(node)
}

func (l *LFUCache) incrFreq(node *Node) {
	f := node.freq
	l.freq[l.minFreq].remove(node)
	if f == l.minFreq && l.freq[l.minFreq].isEmpty() {
		delete(l.freq, l.minFreq)
		l.minFreq++
	}
	f++
	node.freq++
	if _, ok := l.freq[f]; !ok {
		l.freq[f] = NewDLinkList()
	}
	l.freq[f].pushFront(node)
}

type DLinkList struct {
	head, tail *Node
}

func NewDLinkList() *DLinkList {
	head := &Node{}
	tail := &Node{}
	head.post = tail
	tail.pre = head
	return &DLinkList{
		head: head,
		tail: tail,
	}
}

func (l DLinkList) removeTail() int {
	if l.isEmpty() {
		return -1
	}
	node := l.tail.pre
	temp := node.pre
	temp.post = l.tail
	l.tail.pre = temp
	return node.key
}

func (l DLinkList) pushFront(node *Node) {
	temp := l.head.post
	l.head.post = node
	temp.pre = node
	node.pre = l.head
	node.post = temp
}

func (l DLinkList) remove(node *Node) {
	nl, nr := node.pre, node.post
	nl.post = nr
	nr.pre = nl
}

func (l DLinkList) isEmpty() bool {
	return l.head.post == l.tail
}

type Node struct {
	pre, post      *Node
	key, val, freq int
}

func LFU2(operators [][]int, k int) []int {
	// write code here
	l := New(k)
	ans := make([]int, 0)
	for _, op := range operators {
		if op[0] == 1 {
			l.set(op[1], op[2])
		} else {
			ans = append(ans, l.get(op[1]))
		}
	}
	return ans
}
