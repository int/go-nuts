/*
 * Sieve of Eratosthenes in Go
 *
 * some minor optimization, see comments in sieve()
 * each number cost 1 bit in isprime memory
 *
 * fast enough
 *
 * some bit manipulation, time enclapsed, etc.
 *
 *
 * SideNote:
 *
 * if you want super fast prime generator, check out
 * http://cr.yp.to/primegen.html
 * which uses a different (more advanced algorithm): 'Sieve of Atkin'
 * ref: http://cr.yp.to/papers/primesieves.pdf
 * It's very very fast and uses much less memory
 */
package main

import (
	"math"
	"time"
	"fmt"
)

var isprime, prime []uint32

func set(p []uint32, k uint32) { p[k/32] |= 1 << (k & 31) }

func clear(p []uint32, k uint32) { p[k/32] &^= 1 << (k & 31) }

func bit(p []uint32, k uint32) uint32 { return (p[k/32] >> (k & 31)) & 1 }

func sieve(max uint32) int {
	isprime = make([]uint32, max/32+1) // init 0
	f := float64(max)
	prime = make([]uint32, int(2*f/math.Log(f)))
	// init:
	// 2 is the only even prime; mark all odd is prime
	set(isprime, 2)
	prime[0] = 2
	for i := uint32(3); i < max; i += 2 {
		set(isprime, i)
	}
	r := 1
	limit := uint32(math.Ceil(math.Sqrt(f)))
	// try odds
	for i := uint32(3); i < max; i += 2 {
		if bit(isprime, i) == 1 {
			prime[r] = i
			r++
			// un-optimized version might be like:
			// for j := i; j < max; j += i { ... }
			if i > limit {
				continue
			}
			for j := i * i; j < max; j += 2 * i {
				clear(isprime, j)
			}
		}
	}
	return r
}

func main() {
	t := time.Nanoseconds()
	n := uint32(1e8)
	fmt.Printf("primepi(%d): %d\n", n, sieve(n))
	fmt.Printf("takes %.3f seconds\n", float64(time.Nanoseconds()-t)/1e9)
}
