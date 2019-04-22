package main

import (
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	Word  string
	Count int
}

func main() {
	strInput := "I am a student a boy, I am a student a boy, I am a student a boy, I am a student a boy"
	//fmt.Println(strInput[0:3])
	strInput = strings.Replace(strInput, "boy,", "boy", -1)
	// fmt.Println(strInput)

	wordRev := make(map[string]int)
	_ = wordRev

	// preIndex, curIndex := 0, 0
	// for i := 0; i < len(strInput); i++ {
	// 	if strInput[i] == ' ' {
	// 		curIndex = i
	// 		fmt.Println(strInput[preIndex:curIndex])
	// 		preIndex = curIndex + 1
	// 	}
	// }

	// for i, v := range strInput {
	// 	//fmt.Println(i, string(v))
	// 	if v == ' ' || v == '\n' {
	// 		curIndex = i
	// 		fmt.Println(strInput[preIndex:curIndex])
	// 		preIndex = curIndex + 1
	// 	}
	// 	// if v == ' ' || v == '\n' {
	// 	// 	//fmt.Println(i, string(v))

	// 	// 	fmt.Println(strInput[preIndex:i])
	// 	// 	preIndex = i + 1
	// 	// }
	// }

	a := strings.Split(strInput, " ")
	for _, v := range a {
		//fmt.Println(i, v)
		wordRev[v]++
	}

	var keys []string
	for k, _ := range wordRev {
		//fmt.Println(k, v)
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		//fmt.Println(k)
		fmt.Println(k, wordRev[k])
	}
	// wordRev1 := make(map[string]int)
	// for k, v := range wordRev {
	// 	if k ==
	// }
}
