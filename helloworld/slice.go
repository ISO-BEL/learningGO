package main
import "fmt"
//slices are like array references
//go refers to array literals as slice literals
// -- cap() -- len()

func Simple(){
	//var x = []int{1, 3, 55, 2}
	//y := x[2]

	var pool = [300]int{}
	a := pool[:20]  //should exclude the final index n-1
	b := pool[20:40]
	c := pool[40:60]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	slice_print(a)

	//make slice (alloc 0'd) (first num len(), cap() ) 
	k := make([]int, 0, 20)
	fmt.Println(k)

	//slices of slices
	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}

	array2d := [][]int{
		[]int{1, 4, 5, 3, 1, 3},
		[]int{2, 6, 3, 6, 2, 8},
		[]int{3, 8, 9, 5, 0, 1},
	}
	x := array2d[0][0]
	fmt.Println(x)

	//apending a slice
}
func slice_print(a []int){
	fmt.Printf("len = %d, cap = %d, (%v)", len(a), cap(a), a)
}


