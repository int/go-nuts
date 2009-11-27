// play with gorouintes inspired by the Day3 pdf IsReady exmaple
//
// parallel cooking
//
// one cook does not block another
// time ./cook roughly equals to the longest single cook
// instead of sum of the cooks

package main

import "time"

type Cook struct {
	count	int;
	c	chan string;
}

func newCook() *Cook {
	ck := new(Cook);
	ck.c = make(chan string);
	return ck;
}

func (ck *Cook) cook(name string, sec int64) {
	ck.count++;
	go func() {
		println("cooking", name);
		time.Sleep(sec * 1e9);
		ck.c <- name;
	}();
}

func (ck *Cook) wait() {
	for i := 0; i < ck.count; i++ {
		println(<-ck.c, "is ready.")
	}
}

func main() {
	ck := newCook();
	ck.cook("coffee", 5);
	ck.cook("tea", 3);
	ck.cook("cake", 10);
	ck.cook("pie", 8);
	ck.cook("nuts", 2);
	ck.wait();
}
