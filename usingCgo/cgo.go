package main
//
//int function(void)
//{
//	return 5;
//}
import "C"
import "fmt"

func main(){
	x := C.function()
	fmt.Printf("x = %d\n", x)
}
