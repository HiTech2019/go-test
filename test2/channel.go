// Copyright 2018 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	_ "net/http/pprof"
)

//go:generate cp gen.go /tmp/chan_gen.go
//go:generate go run /tmp/chan_gen.go
//go:generate go fmt

// nodeT is channel message
type nodeT struct {
	value interface{} // the message value. i hope it's a happy one
	next  *nodeT      // next item in the queue
}

// Chan represents a single-producer / single-consumer channel.
type Chan struct {
	waitg sync.WaitGroup // used for sleeping. gotta get our zzzs
	queue *nodeT         // items in the sender queue
	recvd *nodeT         // receive queue, receiver-only
	sleep nodeT          // resuable indicates the receiver is sleeping
}

// Send sends a message of the receiver.
func (ch *Chan) Send(value interface{}) {
	n := &nodeT{value: value}
	var wake bool
	for {
		n.next = (*nodeT)(atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
		))
		if n.next == &ch.sleep {
			// there's a sleep placeholder in the sender queue.
			// clear it and prepare to wake the receiver.
			if atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
				unsafe.Pointer(n.next), unsafe.Pointer(n.next.next)) {
				// wake up the receiver
				wake = true
			}
		} else {
			if atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
				unsafe.Pointer(n.next), unsafe.Pointer(n)) {
				break
			}
			/*
				node.next = ch.queue

				if node.next == ch.queue {
					ch.queue = node
					break
				}


				node.next == ch.queue == ch.sleep {
					ch.queue = nil
				}
			*/
		}
		runtime.Gosched()
	}
	if wake {
		ch.waitg.Done()
	}
}

// Recv receives the next message.
func (ch *Chan) Recv() interface{} {
	for {
		if ch.recvd != nil {
			// new message, fist pump
			value := ch.recvd.value
			ch.recvd = ch.recvd.next
			return value
		}
		// let's load more messages from the sender queue.
		for {
			queue := (*nodeT)(atomic.LoadPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
			))
			if queue == nil {
				// sender queue is empty. put the receiver to sleep
				ch.waitg.Add(1)
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
					unsafe.Pointer(queue), unsafe.Pointer(&ch.sleep)) {
					ch.waitg.Wait()
				} else {
					ch.waitg.Done()
				}
			} else if atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&ch.queue)),
				unsafe.Pointer(queue), nil) {
				// we have an isolated queue of messages
				// reverse the queue and fill the recvd list
				for queue != nil {
					next := queue.next
					queue.next = ch.recvd
					ch.recvd = queue
					queue = next
				}
				break
			}
			runtime.Gosched()
		}
	}
}

func main() {

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			num := strconv.FormatInt(int64(runtime.NumGoroutine()), 10)
			w.Write([]byte(num))
		})
		http.ListenAndServe("127.0.0.1:6061", mux) //开启6061端口可以查看golang当前的routine数 查看是否是routine泄露

	}()

	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	var ch Chan
	var wg sync.WaitGroup
	//var val string

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(m int) {
			for n := 0; n < 10000000; n++ {
				//val = fmt.Sprintf("ping-%d-%d\n", m, n)
				//val := strings.Join([]string{"ping", strconv.Itoa(n)}, "")
				var val bytes.Buffer
				val.WriteString(strconv.Itoa(m))
				val.WriteString("-ping-")
				val.WriteString(strconv.Itoa(n))
				ch.Send(val.String())

				//ch.Send(val)
				time.Sleep(time.Millisecond * 5)
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		ch.Send("close")
	}()

	for {
		v := ch.Recv()
		if v != "close" {
			fmt.Println(v)
		} else {
			break
		}

	}

}

/*

top -p `pidof mainpprof`

https://www.jishuwen.com/d/2idk



curl http://localhost:6060/debug/pprof/heap > heap.0.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.1.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.2.pprof
sleep 30
curl http://localhost:6060/debug/pprof/heap > heap.3.pprof


go tool pprof heap.4.pprof

*/
