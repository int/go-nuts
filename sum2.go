/*
   Goroutines are multiplexed as needed onto system threads. When a goroutine
   executes a blocking system call, no other goroutine is blocked.  We will do
   the same for CPU-bound goroutines at some point, but for now, if you want
   user-level parallelism you must set $GOMAXPROCS. or call
   runtime.GOMAXPROCS(n).
*/

// time ./sum2 -- roughly half of ./sum on a SMP (dual-core)
package main

import "runtime"

var parts = 2;

func psum(n int) (r int64) {
	c := make(chan int64);
	g := func(start int, end int) {
		var r int64;
		for i := start; i < end; i++ {
			r += int64(i);
		}
		c <- r;
	};
	for i := 0; i < parts; i++ {
		go g(i * n / parts, (i+1) * n / parts);
	}
	for i := 0; i < parts; i++ {
		r += <-c;
	}
	return r;
}

func main() {
	runtime.GOMAXPROCS(parts);
	println(psum(1000000000));
}
