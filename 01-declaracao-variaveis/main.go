package main

import "fmt"

var x int
var y string
var z bool

var pl = fmt.Println

func main() {
	// A linha abaixo imprime o "valor zero" de cada tipo. 
	pl(x, y, z)
}
