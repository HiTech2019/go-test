package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value := "2019-03-04 13:43:58-1551678238-1"
	scanTempKey := strings.Split(value, " ")
	scanKey := strings.Split(scanTempKey[1], "-")
	tag := strings.Join([]string{scanKey[1], scanKey[2]}, "-")

	timeSecond, _ := strconv.ParseInt(scanKey[1], 10, 64)
	fmt.Println(scanKey)
	fmt.Println(timeSecond)
	fmt.Println(tag)
}
