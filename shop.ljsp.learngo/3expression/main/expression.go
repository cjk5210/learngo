package main

import (
	"errors"
	"log"
)

func check(x int)error{
	if x<=0{
		return errors.New("x<=0")
	}
	return nil
}
func ifControl(){
	x:=10
	if err:=check(x);err==nil{
		x++
		println(x)
	}else{
		log.Fatalln(err)
	}
}
func switchControl(){
	switch x:=5;x{
	default:
		x+=100
		println(x)
	case 6,5:
		x+=50
		println(x)
		if(x==5){
			break
		}
		fallthrough
	case 3,2:
		println(x)
	}
}
func count() int{
	print("count.")
	return 3
}
func forControl(){
	for i,c:=0,count();i<c;i++{
		println("a",i)
	}
	c:=0
	for c<count(){
		println("b",c)
		c++
	}
	data2:=[3]string{"a","b","c"}
	for i,s:=range data2{
		println(i,s)
	}
}
func gotoContinueBreakControl(){
	outer:
		for x:=0;x<5;x++{
			for y:=0;y<10;y++{
				if y>2{
					println()
					continue outer
				}
				if x>2{
					break outer
				}
				print(x,":",y," ")
			}
		}
}
//show
//0: 0 0: 1 0: 2
//1: 0 1: 1 1: 2
//2: 0 2: 1 2: 2


func main() {
	//ifControl()
	//switchControl()
	//forControl()
	gotoContinueBreakControl()






}
