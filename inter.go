// test interfaces and methods
// * function as reciever of method
// * embedding

package main

type Interface interface {
	Foo(int) int
}

type Int int

func (p Int) Foo(a int) int { return int(p) + a }

type Struct struct {
	Int
}

type Func func(byte)

func (f Func) Foo(a int) int {
	f('x')
	return a * 2
}

func pc(c byte) { println(c) }

func main() {
	s := Struct{7}
	println((&s).Foo(8)) // value receiver works for pointer
	println(Func(pc).Foo(3))
}
