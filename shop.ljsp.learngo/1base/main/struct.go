package main

import "fmt"

type user struct {
	name string
	age byte

}

type manager struct {
	user
	tittle string
}

func structWithAnotherStruct(){
	var m manager
	m.name="Tom"
	m.age=29
	m.tittle="CTO"
	fmt.Println(m)
	//will print:{{Tom 29} CTO}
}
func main()  {
	structWithAnotherStruct()
}