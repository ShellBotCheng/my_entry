package lru

import (
	"container/list"
	"errors"
	"sync"
)

var Client *Lru

type Lru struct {
	Max   int
	l     *list.List
	Call  func(key interface{}, value interface{})
	cache map[interface{}]*list.Element
	mu    *sync.Mutex
}

type Node struct {
	Key interface{}
	Val interface{}
}

func InitLru(len int) *Lru {
	Client = &Lru{
		Max:   len,
		l:     list.New(),
		cache: make(map[interface{}]*list.Element),
		mu:    new(sync.Mutex),
	}
	return Client
}

func (l *Lru) Add(key interface{}, val interface{}) error {
	if l.l == nil {
		return errors.New("not init NewLru")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.cache[key]; ok { //以及存在
		e.Value.(*Node).Val = val
		l.l.MoveToFront(e)
		return nil
	}
	ele := l.l.PushFront(&Node{
		Key: key,
		Val: val,
	})
	l.cache[key] = ele
	if l.Max != 0 && l.l.Len() > l.Max {
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache, node.Key)
			if l.Call != nil {
				l.Call(node.Key, node.Val)
			}
		}
	}
	return nil
}

func (l *Lru) Get(key interface{}) (val interface{}, ok bool) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if ele, ok := l.cache[key]; ok {
		l.l.MoveToFront(ele)
		return ele.Value.(*Node).Val, true
	}
	return
}

func (l *Lru) GetAll() []*Node {
	l.mu.Lock()
	defer l.mu.Unlock()
	var data []*Node
	for _, v := range l.cache {
		data = append(data, v.Value.(*Node))
	}
	return data
}

func (l *Lru) Del(key interface{}) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if ele, ok := l.cache[key]; ok {
		delete(l.cache, ele)
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			delete(l.cache, key)
			if l.Call != nil {
				node := e.Value.(*Node)
				l.Call(node.Key, node.Val)
			}
		}
	}

}
