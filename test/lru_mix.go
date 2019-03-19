package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Pair struct {
	key     string
	value   interface{}
	curtime int64
	elem    *list.Element
}

type LRUCache struct {
	capacity int
	m        map[string]*Pair
	list     *list.List
	isExpLru bool
	expire   int64
	lock     sync.RWMutex
}

func Constructor(capacity int, isexp bool, expire int64) *LRUCache {
	lru := new(LRUCache)
	lru.capacity = capacity
	lru.m = make(map[string]*Pair)
	lru.expire = expire
	lru.isExpLru = isexp
	lru.list = list.New()
	return lru
}

func (this *LRUCache) Get(key string) interface{} {
	if this.isExpLru {
		this.lock.Lock()
		pair, ok := this.m[key]
		if ok {
			if time.Now().Unix()-pair.curtime >= this.expire { //时间过期则删除
				this.list.Remove(pair.elem)
				delete(this.m, key)
			} else {
				this.list.MoveToFront(pair.elem)
				this.lock.Unlock()
				return pair.value
			}
		}
		this.lock.Unlock()
		return nil
	} else {
		this.lock.Lock()
		pair, ok := this.m[key]
		if ok {
			this.list.MoveToFront(pair.elem)
			this.lock.Unlock()
			return pair.value
		}
		this.lock.Unlock()
		return nil
	}

}

func (this *LRUCache) Put(key string, value interface{}) *Pair {
	if this.isExpLru {
		this.lock.Lock()
		pair, ok := this.m[key]
		if ok {
			pair.value = value
			pair.curtime = time.Now().Unix()
			this.list.MoveToFront(pair.elem)
			this.lock.Unlock()
			return nil
		} else {
			if len(this.m) >= this.capacity {
				elem := this.list.Back()
				this.list.Remove(elem)
				if pair, ok := elem.Value.(*Pair); ok {
					delete(this.m, pair.key)
				}
			} else {
				// 逆序遍历
				count := 0
				for elem := this.list.Back(); elem != nil; elem = elem.Prev() {
					if pair, ok := elem.Value.(*Pair); ok {
						if time.Now().Unix()-pair.curtime >= this.expire { //时间过期则删除
							this.list.Remove(pair.elem)
							delete(this.m, key)
							break
						} else {
							if count++; count >= 3 {
								break
							}
						}
					}
				}
			}

			pair := &Pair{key: key, value: value, curtime: time.Now().Unix()}
			pair.elem = this.list.PushFront(pair)
			this.m[key] = pair

		}
		this.lock.Unlock()
		return pair
	} else {
		this.lock.Lock()
		pair, ok := this.m[key]
		if ok {
			pair.value = value
			this.list.MoveToFront(pair.elem)
			this.lock.Unlock()
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
		this.lock.Unlock()
		return pair
	}
}

const (
	MaxWorker int = 1000
)

func main() {
	lru := Constructor(100, false, 6)

	var wg sync.WaitGroup
	wg.Add(MaxWorker * 12)

	var i int
	for i = 0; i < MaxWorker*10; i++ {
		go func(n int) {
			for m := 0; m < (n+1)*100000; m++ {
				lru.Put(strconv.Itoa(n), n)
			}

			wg.Done()
		}(i)
	}

	for i = 0; i < MaxWorker*2; i++ {
		go func(n int) {

			for {
				val := lru.Get(strconv.Itoa(n))
				if val != nil {
					fmt.Printf("%d %d\n", n, val.(int))
				} else {
					fmt.Printf("%s", "expired")
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
