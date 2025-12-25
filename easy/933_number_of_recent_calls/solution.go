package main

import "fmt"

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{q: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	winStart := t - 3000
	i := 0
	for i < len(this.q) && this.q[i] < winStart {
		i++
	}
	this.q = this.q[i:]
	return len(this.q)
}

func main() {
	obj := Constructor()
	obj.Ping(1)
	obj.Ping(100)
	obj.Ping(3001)
	param1 := obj.Ping(3002)
	fmt.Println(param1)
}
