/* hello world */
/* namespaces kept seperate in the module syntax */
package main
import "fmt" /*fmt. prefix*/
import "rsc.io/quote"

func main() {
	var sum int
    fmt.Println(quote.Go())
	for i := 0; i < 10; i++{
		sum += i
	}
	Simple()
}
