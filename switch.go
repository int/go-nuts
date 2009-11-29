// switch in Go is more convenient, powerful and flexible
//
// 1. can be a shortened a if-else-if-else chain
// 2. no fall thru but support comma-sep'ed lists.

package main

// stupid but just to demo the use of switch
func even(n int) bool {
	switch n {
	case 2,8,10:
		return true;
	case 3,5,7:
		return false;
	}
	return even2(n);
}

// stupid 2
func even2(n int) bool {
	switch {
	case n == 0:
		return true;
	case n == 1:
		return false;
	case n % 2 == 0:
		return true;
	}
	return false;
}

func main() {
	for i := 0; i <= 10; i++ {
		println(i, even(i));
	}
}
