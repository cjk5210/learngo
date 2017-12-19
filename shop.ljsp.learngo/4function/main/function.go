package main

import (
	"fmt"
	"log"
)

func hello(){
	println("hello,world!")

}
func exec(f func()){//use the function as parameter
	f()
}



type FormatFunc func(string,... interface{})(string,error)

var defineFunctionType FormatFunc
var defineFormat FormatFunc


func useSecondLevelPointerForReturnValue(p **int){
	x:=100
	*p=&x
}

func testUseSecondLevelPointerForReturnValue(){
	var p *int
	useSecondLevelPointerForReturnValue(&p)
	println(*p)
}

func nameTheReturnValue(sql string,index int)(count int,pages int,err error){
	count=1
	pages=2
	err=nil
	return count,pages,err  //There can use return.Because system can return 'count,pages,err'
}
func unNameFunction(){
	func(s string){
		println(s)
	}("hello,world!again!")
}

func anonymousFunctionToVariable(){
	add:=func(x,y int) int{
		return x+y
	}
	println(add(1,4))
}

func anonymousFunctionAsAParameter(f func()){
	f()
}
func testUpFunction(){
	anonymousFunctionAsAParameter(func(){
		println("hello,world!3")
	})
}
func deferTest(){
	defer func() {//be use to close source,file,connect and error handling
		println("print one")

	}()
	defer func(){
		println("print two")
	}()

	println("print three")//will show three two on,the order is unbelievable but fun
}

type DivError struct{
	x,y int
}
func (DivError) Error()string{
	return "division by zero"
}
func div(x,y int)(int,error){
	if y==0{
		return 0,DivError{x,y}//返回自定义错误类型
	}
	return x/y,nil
}
func errMain(){
	z,err:=div(5,0)
	if err!=nil{
		switch e:=err.(type) {
		case DivError:
			fmt.Println(e,e.x,e.y)
		default:
			fmt.Println(e)

		}
		log.Fatalln(err)
	}
	println(z)
}
func main() {
	//f:=hello
	//exec(f)
	//testUseSecondLevelPointerForReturnValue()
	//unNameFunction()
	//anonymousFunctionToVariable()
	//testUpFunction()
	//deferTest()
	errMain()





}
