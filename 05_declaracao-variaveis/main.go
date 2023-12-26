package main
import "fmt"

type age int
var x age
var y int

func main() {
	fmt.Printf("%v %T", x, x)
	x = 42
	fmt.Printf("\n%v %T", x, x)
	y = int(x)
	fmt.Printf("\n%v %T", y, y)
}