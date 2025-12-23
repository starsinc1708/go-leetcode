package main

import (
	"fmt"
	"go-leetcode/myhashmap"
)

func main() {
	myMap := myhashmap.Constructor()
	myMap.Put(1, 10)
	myMap.Put(11, 20)
	fmt.Println(myMap.Get(1))
	fmt.Println(myMap.Get(11))
	myMap.Remove(1)
	fmt.Println(myMap.Get(1))
}
