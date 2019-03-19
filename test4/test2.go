package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 检查指针对象的零值为 nil
	// var p *int
	// fmt.Println("The zero value of a pointer is: ", p)

	// // 指向指针的指针
	// pp := &p
	// fmt.Printf("The type of a pointer points another pointer is: %T\n", pp)

	// // 指针对象赋值
	// intVar := 100000000
	// p = &intVar
	// fmt.Println("After assignment, p is: ", &intVar)
	// fmt.Println("After assignment, p is: ", p)
	// fmt.Println("The value pointer p points is: ", *p)

	// // 使用 unsafe.Pointer 方法可以将一个类型的指针转化为 Pointer
	// // Pointer 可以被转化为任意类型的指针。
	// // 注意由于 int 为 int32 的别名，占 4 个字节，所以我们将其转化为含有 4 个字节元素的 `byte` 数组指针
	// var strP *[4]byte
	// strP = (*[4]byte)(unsafe.Pointer(p))

	// fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer p addr is: \n", p)
	// fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer p addr is: \n", unsafe.Pointer(p))

	// fmt.Printf("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer strP addr is: %p\n", strP)

	// fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer strP is: ", strP)
	// fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer strP points to: ", *strP)

	// // 指针指向的对象内容使用 `.` 而不是 `->` 来进行访问
	// type User struct {
	// 	name string
	// }
	// userP := &User{
	// 	"Xiaohui",
	// }
	// fmt.Println("Before change, The value userP points to is: ", *userP)
	// userP.name = "Ross"
	// fmt.Println("After change,  The value userP points to is: ", *userP)

	//test2
	// intVar := 100000000
	// p := &intVar
	// fmt.Println("intVar is: ", intVar)
	// fmt.Println("intVar addr is: ", &intVar)
	// fmt.Println("p addr is: ", p)
	// fmt.Println("p' addr is: ", &p)

	// fmt.Println("intVar unsafe is: ", unsafe.Pointer(&intVar))
	// fmt.Println("intVar unsafe is: ", (*unsafe.Pointer)(unsafe.Pointer(&intVar)))
	// //fmt.Println("intVar unsafe is: ", &(unsafe.Pointer(&intVar))) //cannot take the address of unsafe.Pointer(&intVar)

	// //test3
	// var n int64 = 5
	// var pn = &n
	// var pf = (*float64)(unsafe.Pointer(pn))
	// // now, pn and pf are pointing at the same memory address
	// fmt.Println(*pf) // 2.5e-323
	// *pf = 3.14159
	// fmt.Println(n) // 4614256650576692846

	//test4
	var n = []byte("hello world...")
	var pf = (*string)(unsafe.Pointer(&n))
	*pf = "my hahaha"
	fmt.Println(string(n))
	fmt.Printf("%p\n", &n)
	fmt.Printf("%p\n", pf)
	fmt.Println(pf)
}
