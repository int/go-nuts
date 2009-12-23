/*
 * segmented sieve:

 * first sieve out small primes <= sqrt(maxlimit) and then
 * do the sieve on each segment.
 *
 *
 * faster than the plain sieve (by a factor of 3~4)
 * also uses less (fixed) memory
 *
 * it's eaiser to get any primes in a range.
 * and it's easier to adapt it to a paralled version
 */

package main

import (
	"math"
	"time"
	"fmt"
)

var (
	sp     []int // small primes
	primes []int // all primes
	pcount int   // primes count
)


// init small primes <= limit
func initsp(limit int) int {
	notprime := make([]bool, limit+1)
	sp = make([]int, limit+1)
	sp[0] = 2
	r := 1
	for i := 3; i <= limit; i += 2 {
		if !notprime[i] {
			sp[r] = i
			r++
			for j := i + i; j <= limit; j += i {
				notprime[j] = true
			}
		}
	}
	sp = sp[0:r]
	return r
}

// sieve primes in [start..end) in one seg
func sieve(start, end int) int {
	if end <= start {
		return 0
	}
	length := end - start
	notprime := make([]bool, length)

	to := int(math.Floor(math.Sqrt(float64(end))))
	for _, p := range sp {
		if p > to {
			break
		}
		// start+i is the first number which is > p and a multiple of p
		i := (-start%p + p) % p
		if start+i == p {
			i += p
		}
		for ; i < length; i += p {
			notprime[i] = true
		}
	}

	r := 0
	for i, b := range (notprime) {
		if !b {
			r++
			//println(start + i)
			primes[pcount] = start + i
			pcount++
		}
	}
	return r

}

// sieve primes in [start..limit]
func seg_sieve(start, limit int) int {
	if start < 2 {
		start = 2
	}
	if start > limit {
		return 0
	}

	limit++
	initsp(int(math.Ceil(math.Sqrt(float64(limit)))) + 4)
	length := int(40000)
	f := float64(limit - start + 8)
	primes = make([]int, int(1.5*f/math.Log(f)+8))

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
	n := int(1e8)
	fmt.Printf("primepi(%d): %d\n", n, seg_sieve(1, n))
	fmt.Printf("takes %.3f seconds\n", float64(time.Nanoseconds()-t)/1e9)
}
