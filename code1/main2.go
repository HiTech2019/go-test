package main

import (
	"fmt"
	"sync"
	"time"
)

type LimitRate struct {
	Rate  int
	Begin time.Time
	Count int
	Lock  sync.Mutex
}

// 1s 达到速率返回true 没有返回false
func (l *LimitRate) Limit() bool {
	var result bool = true
	l.Lock.Lock()

	if l.Count == l.Rate {
		if time.Now().Sub(l.Begin) >= time.Second {
			l.Begin = time.Now()
			l.Count = 0
			result = true
		} else {
			result = false
		}
	} else {
		l.Count++
	}

	l.Lock.Unlock()
	return result
}

func (l *LimitRate) SetRate(r int) {
	l.Rate = r
	l.Begin = time.Now()

}

func (l *LimitRate) GetRate() int {
	return l.Rate
}

func main() {
	var wg sync.WaitGroup
	var lr LimitRate
	lr.SetRate(5)

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			//fmt.Println("Got it...")

			if lr.Limit() {
				fmt.Println("Got it...")
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
