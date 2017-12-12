package main

import (
	"errors"
	"fmt"
)

func ifControl()  {
	x:=100
	if x>0{
		println("x")
	}else if x<0{
		println("-x")
	}else{
		println("0")
	}
}

func switchControl(){
	x:=100
	switch  {
		case x>0:
			println("x")
		case x<0:
			println("-x")
		default:
			println("0")
	}
}
func forLoop()  {
	for i:=0;i<5;i++{
		println(i)
	}
	for i:=4;i>=0;i--{
		println(i)
	}
}
func forLikeWhileLoop(){
	x:=0
	for x<5{
		println(x)
		x++
	}
}
func forLikeWhileTrueQuitWithBreak(){
	x:=4
	for{
		println(x)
		x--
		if x<0{
			break
		}
	}
}
func forUseRange(){
	x:=[] int{100,101,102,103,104}
	for i,n:=range x{
		println(i,":",n)
	}
}
func div(a,b int) (int,error){
	if b==0{
		return 0,errors.New("division by zero")
	}
	return a/b,nil
}
func multiReturnValue(){
	a,b:=10,2
	c,err:=div(a,b)
	fmt.Println(c,err)
}
func funcAsAReciveOrReturnValue(x int) func(){
	return func(){
		println(x)
	}
}
func useDeferStartAFuncBeforeFinish(a,b int){
	defer println("Releases some Resources")
	defer println("Resources has been released!")//if a func contains multi-defer ,they will be according to the order like " first in last out" for executed.
	println(a/b)
}
func automaticExpanseArray(){
	x:=make([] int,0,5)
	for i:=0;i<8;i++{
		x=append(x,i)
	}
	fmt.Println(x)
}
func mapIsContainsAKey(){
	m:=make(map[string]int)
	m["a"]=1
	x,ok:=m["b"]
	fmt.Println(x,ok)
}
func main()  {
	//ifControl()
	//switchControl()
	//forLoop()
	//forLikeWhileLoop()
	//forLikeWhileTrueQuitWithBreak()
	//forUseRange()
	//multiReturnValue()

	//f:=funcAsAReciveOrReturnValue(101)
	//f()

	//useDeferStartAFuncBeforeFinish(10,2)
	//automaticExpanseArray()
	mapIsContainsAKey()
}
