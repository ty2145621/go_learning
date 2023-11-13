package real

import (
	"container/list"
)

/*
*
BM100 设计LRU缓存结构
*/
type kv struct {
	k int
	v int
}

type Solution struct {
	l   *list.List
	m   map[int]*list.Element // key => element
	cap int
}

func Constructor(capacity int) Solution {
	s := Solution{
		l:   list.New(),
		m:   make(map[int]*list.Element, capacity),
		cap: capacity,
	}
	return s
}

func (s *Solution) get(key int) int {
	if e, ok := s.m[key]; !ok {
		return -1
	} else {
		s.l.MoveToFront(e)
		return e.Value.(kv).v
	}
}

func (s *Solution) set(key int, value int) {
	e, ok := s.m[key]
	if !ok {
		e = s.l.PushFront(kv{k: key, v: value})
		if s.l.Len() > s.cap {
			oldkv := s.l.Remove(s.l.Back()).(kv)
			delete(s.m, oldkv.k)
		}
		s.m[key] = e
	} else {
		e.Value = kv{k: key, v: value}
		s.l.MoveToFront(e)
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * solution := Constructor(capacity);
 * output := solution.get(key);
 * solution.set(key,value);
 */
