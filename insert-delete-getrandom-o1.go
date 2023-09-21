package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.set[val] {
		return false
	}
	this.set[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.set[val] {
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randNum := rand.Intn(len(this.set))
	count := 0
	for k, _ := range this.set {
		if randNum == count {
			return k
		}
		count++
	}
	return -1
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)
	fmt.Println(param_1)
}
