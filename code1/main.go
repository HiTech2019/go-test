package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const size = 100

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < size; i++ {
		wg.Add(1)
		go watch(ctx, &wg, fmt.Sprintf("goroutine %d", i))
	}

	time.Sleep(600 * time.Second)
	fmt.Println("开始结束goroutine")
	cancel()

	wg.Wait()
}

func watch(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "over")
			return
		default:
			fmt.Println(name, "running")
			time.Sleep(time.Second * 2)
		}
	}

}

// import (
// 	"fmt"
// )

// type Slice []int

// func NewSlice() Slice {
// 	return make(Slice, 0)
// }

// func (s *Slice) Add(elem int) *Slice {
// 	*s = append(*s, elem)
// 	fmt.Print(elem, "sss")
// 	return s
// }

// // func main() {
// // 	// s := NewSlice()
// // 	// defer s.Add(1).Add(2)
// // 	// s.Add(3)
// // 	// str := []byte("hello")

// // 	// str[0] = 'x'

// // 	// fmt.Println(string(str))

// // 	list := make([]int, 0)
// // 	list = append(list, 1)
// // 	fmt.Println(list)
// // }

// // func main() {
// // 	type MyInt1 int
// // 	type MyInt2 = int
// // 	var i int = 9
// // 	var i1 MyInt1 = MyInt1(i)
// // 	var i2 MyInt2 = i
// // 	fmt.Println(i1, i2)

// // }

// type User struct {
// }
// type MyUser1 User
// type MyUser2 = User

// func (MyUser1) m1() {
// 	fmt.Println("MyUser1.m1")
// }
// func (User) m2() {
// 	fmt.Println("User.m2")
// }
// func main() {
// 	var i1 MyUser1
// 	var i2 MyUser2
// 	i1.m1()
// 	i2.m2()
// }
