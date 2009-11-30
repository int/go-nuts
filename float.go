// test go's floating point performace

package main


func main() {
	for x := 0; x < 100000; x++ {
		a := float64(1);
		for i := 200; i < 221; i++ {
			a *= float64(i + x);
		}
	}
}
