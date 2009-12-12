// trying Go's big package for multi-precision arithmetic

package main

import "big"

func main() {
	for x := 0; x < 100000; x++ {
		a := big.NewInt(1)
		for i := 200; i < 221; i++ {
			a.Mul(big.NewInt(int64(i+x)), a)
		}
	}
}
