package main

func sum(n int64) (r int64) {
	var i int64
	for ; i < n; i++ {
		r += i
	}
	return r
}

func main()	{ println(sum(1000000000)) }
