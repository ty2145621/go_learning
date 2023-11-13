package real

import (
	"container/list"
	"fmt"
)

/**
BM101 设计ClassLFU缓存结构
*/

func LFU(operators [][]int, k int) []int {
	var ans []int
	lfu := newClassLFU(k)
	for k, op := range operators {
		fmt.Println("before :", k, lfu)
		switch op[0] {
		case 1:
			lfu.Set(op[1], op[2])
		case 2:
			ans = append(ans, lfu.Get(op[1]))
		}
		fmt.Println("after :", k, lfu)
	}

	return ans
}

type freqNode struct {
	l    *list.List
	freq int
}

type kvNode struct {
	key   int
	value int
	freq  int
}

type ClassLFU struct {
	bl      *list.List            // bucketList 其实bucket不用List，反而更好写代码，用一个minFreq保存最小的freq
	freqMap map[int]*list.Element // freq => freqNode
	nodeMap map[int]*list.Element // key => kvNode
	cap     int
}

func newClassLFU(cap int) *ClassLFU {
	return &ClassLFU{
		bl:      list.New(),
		freqMap: make(map[int]*list.Element, 0),
		nodeMap: make(map[int]*list.Element, 0),
		cap:     cap,
	}
}

func (c *ClassLFU) String() string {
	var str = ""
	for k, v := range c.freqMap {
		str += fmt.Sprintf("%d:%+v len:%d  ", k, v.Value.(*freqNode), v.Value.(*freqNode).l.Len())
	}
	str += fmt.Sprintf("\n len Node:%d  ", len(c.nodeMap))

	for k, v := range c.nodeMap {
		str += fmt.Sprintf("%d:%+v  ", k, v.Value.(kvNode))
	}
	str += fmt.Sprintf("\n")
	return str
}

func (c *ClassLFU) Len() int {
	return len(c.nodeMap)
}

func (c *ClassLFU) Get(key int) int {
	node, ok := c.nodeMap[key]
	if !ok {
		return -1
	}
	// 更新freq
	c.removeFromFreqList(node)

	nodeValue := node.Value.(kvNode)
	nodeValue.freq++
	node.Value = nodeValue

	c.insertIntoFreqList(node)

	return node.Value.(kvNode).value
}

func (c *ClassLFU) Set(key int, value int) {
	node, ok := c.nodeMap[key]
	if !ok {
		newKVNode := kvNode{key: key, value: value, freq: 1}
		c.newNode(true, newKVNode)
		return
	}

	c.removeFromFreqList(node)

	nodeValue := node.Value.(kvNode)
	nodeValue.freq++
	nodeValue.value = value
	node.Value = nodeValue

	c.insertIntoFreqList(node)
}

func (c *ClassLFU) removeFromFreqList(node *list.Element) {
	freqListElm := c.freqMap[node.Value.(kvNode).freq]
	fNode := freqListElm.Value.(*freqNode)
	fNode.l.Remove(node)
	delete(c.nodeMap, node.Value.(kvNode).key)

	if fNode.l.Len() == 0 {
		c.bl.Remove(freqListElm)
		delete(c.freqMap, fNode.freq)
	}
}

func (c *ClassLFU) removeMinFreqList() {
	if c.Len() < c.cap {
		return
	}

	minFreqListEle := c.bl.Front()
	minFreqList := minFreqListEle.Value.(*freqNode)
	kvEle := minFreqList.l.Front()
	minFreqList.l.Remove(kvEle)
	delete(c.nodeMap, kvEle.Value.(kvNode).key)

	if minFreqList.l.Len() == 0 {
		c.bl.Remove(minFreqListEle)
		delete(c.freqMap, minFreqList.freq)
	}
}
func (c *ClassLFU) newFreq(first bool, freq int) *list.Element {
	fNode := &freqNode{
		l:    list.New(),
		freq: freq,
	}
	var e *list.Element
	if first {
		e = c.bl.PushFront(fNode)
	} else {
		if oldFreqEle, ok := c.freqMap[freq-1]; ok {
			e = c.bl.InsertAfter(fNode, oldFreqEle)
		} else {
			e = c.bl.PushFront(fNode)
		}
	}
	c.freqMap[fNode.freq] = e
	return e
}

// 新建一个freqList 和node
func (c *ClassLFU) newNode(first bool, newKVNode kvNode) {
	c.removeMinFreqList()

	var freqEle *list.Element
	freqEle, ok := c.freqMap[newKVNode.freq]
	if !ok {
		freqEle = c.newFreq(first, newKVNode.freq)
	}
	kvEle := freqEle.Value.(*freqNode).l.PushBack(newKVNode)
	c.nodeMap[newKVNode.key] = kvEle
}

// 插入一个node
func (c *ClassLFU) insertIntoFreqList(node *list.Element) {
	freqEle, ok := c.freqMap[node.Value.(kvNode).freq]
	if !ok {
		freqEle = c.newFreq(false, node.Value.(kvNode).freq)
	}
	kvEle := freqEle.Value.(*freqNode).l.PushBack(node.Value.(kvNode))
	c.nodeMap[node.Value.(kvNode).key] = kvEle
}
