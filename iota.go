// test iota:  http://golang.org/doc/go_spec.html#Iota
// the implicit repetition of the last non-empty expression list.

package main

const (
	bit0, mask0	= 1 << iota, 1<<iota - 1	// bit0 == 1, mask0 == 0
	bit1, mask1	// bit1 == 2, mask1 == 1
	_, _		// skips iota == 2
	bit3, mask3	// bit3 == 8, mask3 == 7
)

func main() {
	println(bit0, mask0)
	println(bit1, mask1)
	println(bit3, mask3)
}
