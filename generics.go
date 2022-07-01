package main

import (
	"fmt"
)

func PrintSlice[T any](s []T) {
	fmt.Println("Generics")
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})

}

//numeric
/*
package main

import (
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
	//~int //anything derived from int
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("4 + 3 =", Add(4, 3))
	fmt.Println("4.1 + 3.2 =", Add(4.1, 3.2))
}
*/
