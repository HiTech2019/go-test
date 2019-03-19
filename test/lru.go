package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
)

type Pair struct {
	key   string
	value interface{}
	elem  *list.Element
}

type LRUCache struct {
	capacity int
	m        map[string]*Pair
	list     *list.List
	lock     sync.RWMutex
}

func Constructor(capacity int) *LRUCache {
	lru := new(LRUCache)
	lru.capacity = capacity
	lru.m = make(map[string]*Pair)
	lru.list = list.New()
	return lru
}

func (this *LRUCache) Get(key string) interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	pair, ok := this.m[key]
	if ok {
		this.list.MoveToFront(pair.elem)
		return pair.value
	}
	return nil
}

func (this *LRUCache) Put(key string, value interface{}) *Pair {
	this.lock.Lock()
	defer this.lock.Unlock()
	pair, ok := this.m[key]
	if ok {
		pair.value = value
		this.list.MoveToFront(pair.elem)
		return nil
	} else {
		if len(this.m) >= this.capacity {
			elem := this.list.Back()
			this.list.Remove(elem)
			if pair, ok := elem.Value.(*Pair); ok {
				delete(this.m, pair.key)
			}
		}
		pair := &Pair{key: key, value: value}
		pair.elem = this.list.PushFront(pair)
		this.m[key] = pair

	}
	return pair
}

const (
	MaxWorker int = 1000000
)

func main() {
	lru := Constructor(100)

	var wg sync.WaitGroup
	wg.Add(MaxWorker * 2)

	var i int
	for i = 0; i < MaxWorker; i++ {
		go func(n int) {
			lru.Put(strconv.Itoa(n), n)
			wg.Done()
		}(i)
	}

	for i = 0; i < MaxWorker; i++ {
		go func(n int) {
			val := lru.Get(strconv.Itoa(n))
			if val != nil {
				fmt.Printf("%d %d\n", n, val.(int))
			} else {

			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
