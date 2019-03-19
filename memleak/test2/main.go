package main

import "fmt"

var s0 []int

func g(s1 []int) {
	// Assume the length of s1 is much larger than 30.

	fmt.Printf("s1 %p %v\n", s1, s1)

	// s0 = s1[len(s1)-10:]
	// s1 = nil

	s0 = s1[0 : len(s1)-10]
	s1 = nil
	fmt.Printf("s0 %p %v\n", s0, s0)
	fmt.Printf("s1 %p %v\n", s1, s1)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 65, 67, 7, 7, 34, 34, 8, 8, 324, 8, 8, 343, 42, 22, 54, 35, 3, 4, 5, 3, 4, 6, 54, 64, 567, 56, 768, 234, 23, 4, 2, 4, 9, 8, 9, 9, 23, 74, 2, 9347, 293, 749, 237, 4}

	fmt.Printf("array %p %v\n", array, array)
	g(array)

}

//

// func main() {
// 	str := []byte("this is a fucking string")
// 	str2 := str[0:5]
// 	str3 := str2[0:2]
// 	fmt.Printf("%p\n", str)
// 	fmt.Printf("%p\n", str2)
// 	fmt.Printf("%p\n", str3)
// }

// // 0xc00008e000
// // 0xc00008e000
// // 0xc00008e000
