package myhashmap

type KeyValuePair struct {
	Key   int
	Value int
	Used  bool
}

type MyHashMap struct {
	size  int
	table []KeyValuePair
}

func Constructor() MyHashMap {
	size := 1000
	return MyHashMap{
		table: make([]KeyValuePair, size),
		size:  size,
	}
}

func (this *MyHashMap) hash(x int) int {
	return x % this.size
}

func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	for this.table[index].Used {
		if this.table[index].Key == key {
			this.table[index].Value = value
			return
		}
		index = (index + 1) % len(this.table)
	}
	this.table[index] = KeyValuePair{key, value, true}
}

func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	for this.table[index].Used {
		if this.table[index].Key == key {
			return this.table[index].Value
		}
		index = (index + 1) % len(this.table)
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := this.hash(key)
	for this.table[index].Used {
		if this.table[index].Key == key {
			this.table[index].Used = false
			return
		}
		index = (index + 1) % len(this.table)
	}
}
