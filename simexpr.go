// evaluate simple single expression
//
// demo use of exp/eval
// basic input/output
package main

import (
	"fmt"
	"exp/eval"
	"os"
	"bufio"
)

func main() {
	w := eval.NewWorld()
	r := bufio.NewReader(os.Stdin)
	for {
		s, e := r.ReadString('\n')
		if e == os.EOF {
			break
		}
		c, err := w.Compile(s)
		if err == nil {
			v, e := c.Run()
			if e == nil && v != nil {	// single expression
				fmt.Println(v)
			}
		} else {
			fmt.Println("compile error")
		}
	}
}
