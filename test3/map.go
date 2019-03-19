package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"
)

type SrvHttp struct {
	MemMap map[string][]*Record
	sync.Mutex
}

type Record struct {
	RandCode string
	Second   int64
}

func (server *SrvHttp) RecordWordToMem(openID string, randCode string) {
	// length := 0
	// server.MemMap.Range(func(k, v interface{}) bool {
	// 	length++
	// 	return true
	// })

	//fmt.Println(length)

	elems := &Record{randCode, time.Now().Unix()}

	server.Lock()
	defer server.Unlock()

	recordList, ok := server.MemMap[openID]
	if ok {
		recordList = append(recordList, elems)
		server.MemMap[openID] = recordList
	} else {
		var list []*Record
		list = append(list, elems)
		server.MemMap[openID] = list
		//fmt.Println("add 1")
	}

	// if length := len(server.MemMap); length > 100 {
	// 	delCount := length - 100
	// 	count := 0
	// 	for k, _ := range server.MemMap {
	// 		delete(server.MemMap, k)
	// 		count++
	// 		//fmt.Println("delete", k)
	// 		if count == delCount {
	// 			break
	// 		}
	// 	}
	// }

	// server.MemMap.Range(func(k, v interface{}) bool {
	// 	length++
	// 	return true
	// })

	// fmt.Println(length)

	// if length >= 100 {
	// 	server.MemMap.Delete("1")
	// 	server.MemMap.Range(func(k, v interface{}) bool {
	// 		//fmt.Println(k, v)
	// 		curSecond := time.Now().Unix()
	// 		var newArray []Record
	// 		if array, ok := v.([]Record); ok {
	// 			for i := 0; i < len(array); i++ {
	// 				if curSecond-array[i].Second > 2*60 { //大于两分钟
	// 					newArray = append(newArray, array[:i], array[i+1:])
	// 				}
	// 			}
	// 		}
	// 		return true
	// 	})
	// }

}

func (server *SrvHttp) CleanToLimit() {
	// if length := len(server.MemMap); length > 100 {
	// 	delCount := length - 100
	// 	count := 0
	// 	for k, vlist := range server.MemMap {
	// 		//fmt.Println(k, val)
	// 		//_ = val
	// 		if vlist != nil {
	// 			for i := 0; i < len(vlist); i++ {
	// 				vlist[i] = nil
	// 			}
	// 			vlist = nil
	// 		}

	// 		delete(server.MemMap, k)
	// 		count++
	// 		//fmt.Println("delete", k)
	// 		if count == delCount {
	// 			break
	// 		}
	// 	}

	// }
	for k, vlist := range server.MemMap {
		//fmt.Println(k, val)
		//_ = val

		if vlist != nil {
			for i := 0; i < len(vlist); i++ {
				vlist[i] = nil
				fmt.Println("del")
			}
			vlist = nil
		}
		server.MemMap[k] = nil
		delete(server.MemMap, k)
	}

}
func main() {

	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	//server := new(SrvHttp)
	var server SrvHttp
	server.MemMap = make(map[string][]*Record)

	var wg sync.WaitGroup
	//var val string
	//server.RecordWordToMem("key", "123456")
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(x int) {
			for m := 0; m < 1000; m++ {
				server.RecordWordToMem("key"+strconv.Itoa(x), "123456")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	// for k, v := range server.MemMap {
	// 	fmt.Println(k, len(v))
	// }

	server.CleanToLimit()

	for k, v := range server.MemMap {
		fmt.Println(k, len(v))
	}

	server.MemMap = nil
	//server = nil

	fmt.Println("ok")
	time.Sleep(time.Second * 10000000)
}
