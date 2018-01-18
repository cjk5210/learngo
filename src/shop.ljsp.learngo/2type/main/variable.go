package main

import "fmt"

func defineVariable(){
	var x int		//the default value is 0
	var y=false		//auto-judge the variable type from the init value
	//offical proposal method
	var m,n int		//define multi-variable use the same type
	var a,s=100,"abc" //define multi-variable and give the init difference value
	fmt.Println(x,y,m,n,a,s)
}

func shortDefine(){//CAUTION: 1.定义变量，同时显式初始化 2.不能写上数据类型 3.只能用在函数内部
	x:=100
	a,s:=1,"abc"
	println(x,a,s)
}
var x=100
func globalVarAndLocalVar(){
	println(&x,x)			//global var


	x:="abc"
	println(&x,x)			//local var
}
func theUseageOfShortDefine(){	//use in if/for/switch
	x:=100
	println(&x)
	x,y:=200,"abc"			//CAUTION: at multi-define, There must be have one var haven't be define
	println(&x,x)
	println(y)

}
func defineConstant(){
	const x,y int=123,0x22
	const s="hello,world!"
	const c="我"			//if you don't use this constant ,the system can't show the error.

}

func usePreviousLineType(){
	const(
		x uint16=120
		y			//if you don't give a type to a variable ,it will use the type from previous line
		s="abc"
		z
	)
	fmt.Printf("%T,%v\n",y,y)
	fmt.Printf("%T,%v\n",z,z)
}



func main() {
	//defineVariable()
	//shortDefine()
	//theUseageOfShortDefine()
	usePreviousLineType()






}
