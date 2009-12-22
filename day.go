// slice, method, stringer, typed const, etc..
package main

import "fmt"

type Day int

// := only can be used inside a function
// it cannot be used in top-level declaration
var dayName = []string{"Mon", "Tues", "Wednes", "Thurs", "Fri", "Satur", "Sun"}

func (day Day) String() string { return dayName[day] + "day" }

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	day := Thursday
	fmt.Println(day)
	fmt.Printf("%T\n", Sunday) // print type
}
