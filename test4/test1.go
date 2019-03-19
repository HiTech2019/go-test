package main

import (
	"fmt"
)

// var a, b, c int

// func main() {
// 	go func() {
// 		a = 1
// 		b = 2
// 	}()
// 	go func() {
// 		c = a + 2
// 	}()

// 	time.Sleep(time.Second * 1)
// 	log.Println(a, b, c)
// }

const (
	BindPhone = 1
	PwdWord   = 2
)

func main() {

	InputType := BindPhone
	for {
		switch InputType {
		case BindPhone:
			fmt.Println("绑定数据")
			goto end
		case PwdWord:
			fmt.Println("积分口令")
			goto end
		default:
			fmt.Println("未知口令")
			goto end1
		}

	end:
		break

	end1:
	}

}
