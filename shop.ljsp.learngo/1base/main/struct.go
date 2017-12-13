package main

import (
	"fmt"
	"time"
)

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




type X int
func (x *X) inc(){
	*x++
}
func typeAInt(){
	var x X
	x.inc()
	println(x)
}



type user1 struct {
	name string
	age byte
}
type manager1 struct {
	user1
	title string
}
func anonymousInnerClass()  {
	var m manager1
	m.name="TOM"
	m.age=29
	println(m.ToString)
	println(m.name,"--",m.age)
}
func (u user1)ToString() string  {
	return fmt.Sprintf("%+v",u)
}



type user2 struct{
	name string
	age byte
}
func(u user2) Print(){
	fmt.Printf("%+v",u)
	fmt.Println(u)
}
type Printer interface {
	Print()
}
func interfaceTest(){
	var u user2
	u.name="Tom"
	u.age=29
	var p Printer=u
	p.Print()
}



func task(id int){
	for i:=0;i<5;i++{
		fmt.Printf("%d: %d\n",id,i)
		time.Sleep(time.Second)
	}
}
func testConcurrentExecution(){
	go task(1)
	go task(2)
	time.Sleep(time.Second*6)
}




func main()  {
	//structWithAnotherStruct()
	//typeAInt()
	//anonymousInnerClass()
	//interfaceTest()
	//testConcurrentExecution()






}