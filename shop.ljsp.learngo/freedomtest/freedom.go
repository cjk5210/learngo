package main

import (
	"fmt"
	"reflect"
)

const (
	fr float32  =100
)
var(
	a=32
	b=32.0
	c=true
)

func main() {

	//println(fr,"lic")
	fmt.Printf("type:",reflect.TypeOf(c))

}
