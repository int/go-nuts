/*
 * wheel segmented sieve:
 *
 * as an alternative to segemented sieve (seg_sieve.go),
 * this version does not sieve out small primes up to sqrt(maxlimit),
 * as that sqrt might be too large or you just dont want to
 *
 * refs:
 * http://en.wikipedia.org/wiki/Wheel_factorization
 * http://plan9.bell-labs.com/sources/plan9/sys/src/cmd/primes.c
 *
 */

package main

import (
	"math"
	"time"
	"fmt"
)

var (
	sp     []int   // first several small primes
	dnext  []int   // delta to next possible prime
	primes []int64 // all primes
	pcount int     // primes count
)


// helper: return true if n is not divided by any of primes in sp
func ok(n int) bool {
	for _, p := range sp {
		if n%p == 0 {
			return false
		}
	}
	return true
}


// init wheel sieve: the wheel length can be speicified dynamically
func initwheel(wlen int) {
	sp = []int{2, 3, 5, 7, 11, 13} // first small primes
	if wlen <= 0 {
		wlen = 4 // 2*3*5*7 is a good choice
	} else if wlen > len(sp) {
		wlen = len(sp)
	}
	sp = sp[0:wlen]

	wheel := 1
	for i := 0; i < wlen; i++ {
		wheel *= sp[i]
	}
	candi := make([]int, wheel)
	clen := 0
	for i := 1; i <= wheel; i++ {
		if ok(i) {
			candi[clen] = i
			clen++
		}
	}
	dnext = make([]int, clen)
	for i := 0; i < clen-1; i++ {
		//println(candi[i], candi[i+1])
		dnext[i] = candi[i+1] - candi[i]
	}
	dnext[clen-1] = wheel + 1 - candi[clen-1]
	/*
		for _, p := range dnext {
			println(p)
		}
	*/
}


// helper
func mark(start, length, p int64, notprime []bool) {
	//println("mark", p)
	// start+i is the first number which is > p and a multiple of p
	i := (-start%p + p) % p
	//println("got", i, start+i)
	if start+i == p {
		i += p
	}
	for ; i < length; i += p {
		//println("cross", start+i)
		notprime[i] = true
	}
}

// sieve primes in [start..end) in one seg
func sieve(start, end int64) int {
	if end <= start {
		return 0
	}

	length := end - start
	notprime := make([]bool, length)
	to := int64(math.Floor(math.Sqrt(float64(end))))

	// first few real primes
	for _, p := range sp {
		mark(start, length, int64(p), notprime)
	}

	// candidates on the wheel
	p := int64(1)
	for i := 0; ; i = (i + 1) % len(dnext) {
		p += int64(dnext[i])
		if p > to {
			break
		}
		mark(start, length, p, notprime)
	}

	r := 0
	for i, b := range (notprime) {
		if !b {
			r++
			//println(pcount, start+int64(i))
			primes[pcount] = start + int64(i)
			pcount++
		}
	}
	return r

}

// sieve primes in [start..limit]
func seg_sieve(start, limit int64, wlen int) int {
	if start < 2 {
		start = 2
	}
	if start > limit {
		return 0
	}

	limit++
	initwheel(wlen)
	length := int64(40000)
	f := float64(limit - start + 8)
	primes = make([]int64, int(1.5*f/math.Log(f)+8))
	pcount = 0

	r := 0
	for ; start+length <= limit; start += length {
		r += sieve(start, start+length)
	}
	r += sieve(start, limit)
	primes = primes[0:pcount]
	return r
}

func main() {
	t := time.Nanoseconds()
	n := int64(1e8)
	fmt.Printf("primepi(%d): %d\n", n, seg_sieve(1, n, 0))
	t1 := time.Nanoseconds()
	fmt.Printf("takes %.3f seconds\n", float64(t1-t)/1e9)
	a, b := int64(1e15-1e5), int64(1e15)
	fmt.Printf("# of prims in [%d..%d]: %d\n", a, b, seg_sieve(a, b, 6))
	t2 := time.Nanoseconds()
	fmt.Printf("takes %.3f seconds\n", float64(t2-t1)/1e9)
}
