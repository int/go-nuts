// test go sorting performance
package main

import "sort"

func main() {
	s := make([]int, 1000000);
	for i, _ := range s {
		s[i] = 1000000 - i;
	}
	println(s[999999]);
	sort.SortInts(s);
	println(s[999999]);
}
