package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"
)

//生成32位随机序列
var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	codeLen = int64(len(codes))
)

func createStringWithLengthOnHeap(len int) string {
	data := make([]byte, len)
	for i := 0; i < len; i++ {
		idx, _ := rand.Int(rand.Reader, big.NewInt(codeLen))
		data[i] = byte(codes[idx.Int64()])
	}
	retStr := string(data)
	data = nil
	return retStr
}

var s0 string // a package-level variable

// A demo purpose function.
func f(s1 string) {
	//s0 = s1[:50]
	s0 = (" " + s1[:50])[1:]
	//time.Sleep(time.Second * 1000000)
	// Now, s0 shares the same underlying memory block with s1.
	// Although s1 is not alive now, but s0 is still alive, so
	// the memory block they share couldn't be collected,
	// though there are only 50 bytes used in the block and
	// all other bytes in the block become unavailable.
}

func demo() {
	s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	//fmt.Println(s)

	f(s)
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

	time.Sleep(time.Second * 1)

	// var wg sync.WaitGroup
	// wg.Add(50)

	// for i := 0; i < 50; i++ {
	// 	go func(n int) {
	// 		demo()
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()
	demo()

	fmt.Println("end.................")
	time.Sleep(time.Second * 1000000)
}
