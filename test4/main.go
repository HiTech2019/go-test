package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello world...........lhy")
	// fmt.Println("hello world...........lhy")
	// fmt.Println("hello world...........lhy")
	// fmt.Println("hello world...........lhy")

	// fmt.Println("....................kkkkk")
	// //return

	// ss := fmt.Sprintf("%d", 123123)
	// s := fmt.Sprintf("%s", "sdadasd")
	// fmt.Println(s, ss)
	array := []string{"hello", "Hi", "liuhy"}
	array = append(array, "elems")

	fmt.Println(array)

	var arr interface{} = "['hello' 'Hi' 'liuhy' 'elems']"
	value, flag := arr.([]string)
	if flag {
		fmt.Printf("%T->%d\n", value, flag)
	} else {
		fmt.Println("类型不匹配")
	}
	// var listoflists []string
	// dec := json.NewDecoder(strings.NewReader(arr))
	// err := dec.Decode(listoflists)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(listoflists)
	// }

	// 	strings.s

	// 	arr := "[hello Hi liuhy elems]"
	// 	strArray := strings.Fields(arr)
	// 	for k, i := range strArray {
	// 		fmt.Println(k, i)
	// 	}
}

//大 撒旦爱的阿萨德按时达到啊飒飒

func func1() error {
	fmt.Printf("%s", "hello wokd..........")
	fmt.Print("sssssssssssssssssssssss")
	return nil
}
