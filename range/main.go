package main




import "fmt"

type Foo struct {
	bar string
}

func main() {
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	// for i := 0; i < len(list); i++ {
	// 	fmt.Printf("%p %s\n", &list[i], list[i])
	// }
	list2 := make([]*Foo, len(list))
	//for i, value := range list { //range值copy
	for i, _ := range list {
		//fmt.Printf("%p %v\n", &value, value)
		//list2[i] = &value
		list2[i] = &list[i]
	}

	//fmt.Printf("%p %p %p\n", &list[0], &list[1], &list[2])
	fmt.Println(list[0], list[1], list[2])

	//fmt.Printf("%p %p %p\n", list2[0], list2[1], list2[2])
	fmt.Println(list2[0], list2[1], list2[2])
}
