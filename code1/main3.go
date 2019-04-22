package main

import (
	"fmt"
	"sync"
	"time"
)

type LimitRate struct {
	Rate       int
	Interval   time.Duration
	LastAction time.Time
	Lock       sync.Mutex
}

func (l *LimitRate) Limit() bool {
	result := false
	for {
		l.Lock.Lock()

		if time.Now().Sub(l.LastAction) > l.Interval {
			l.LastAction = time.Now()
			result = true
		}

		l.Lock.Unlock()

		if result {
			return result
		}

		time.Sleep(l.Interval)
	}
}

func (l *LimitRate) SetRate(r int) {
	l.Rate = r
	l.Interval = time.Microsecond * time.Duration(1000*1000/l.Rate)
}

func (l *LimitRate) GetRate() int {
	return l.Rate
}

func main() {
	var wg sync.WaitGroup
	var lr LimitRate
	lr.SetRate(5)

	b := time.Now()
	for i := 0; i < 100; i++ {
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
	fmt.Println(time.Since(b))
}
