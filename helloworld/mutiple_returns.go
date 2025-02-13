package main
import "fmt"

func main(){
	y := 10
	x := 20
	result, error := add(x, y)	
	fmt.Printf("result = %v\n", result)
	fmt.Printf("error = %v\n", error)
}
func add(x int, y int) (result int, error bool) {
	if x == 0 || y == 0 {
		error = true	
	}
	error = false
	result = x + y
	return //naked return
}

