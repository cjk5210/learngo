package main

import (
	"fmt"
	"reflect"
	"unsafe"
	"strings"
	"bytes"
	"unicode/utf8"
)

func stringMain(){
	s:="雨痕\x61\142\u0041"
	fmt.Printf("%s\n",s)
	fmt.Printf("%x,len:%d\n",s,len(s))
	var defaultValue string
	println(defaultValue=="")//true,the init-value is "" for a String variable

	st:=`line\r\n,
		line2`
	println(st)		//use the symbol "`" to define orignal string
}


func sliceString(){
	s:="abcdefg"
	s1:=s[:3]
	s2:=s[1:4]
	s3:=s[2:]
	println(s1,s2,s3)
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func goThroughString(){//byte and rune
	s:="雨痕"
	for i:=0;i<len(s);i++{ //byte
		fmt.Printf("%d:[%c]\n",i,s[i])
	}

	for i,c:=range s{//rune   return index and Unicode symbol
		fmt.Printf("%d:[%c]\n",i,c)

	}
}

func appendJoin(){
	s:=make([] string,1000)
	for i:=0;i<1000;i++{
		s[i]="a"
	}
	b:=strings.Join(s,"")//use strings.Join() append string
	fmt.Println(b)

}

func appendJoinString() string{
	var b bytes.Buffer
	b.Grow(1000)
	for i:=0;i<1000;i++{
		b.WriteString("a")
	}
	return b.String()
}

func testUnicode(){//use Single quotation mark for mark the Unicode string
	r:='我'
	fmt.Printf("%T\n",r)

}
func stringConvert(){
	r:='我'
	s:=string(r) //rune to string
	b:=byte(r)// rune to byte
	s2:=string(b)
	r2:=rune(b)
	fmt.Println(s,b,s2,r2)
}

func lenUtf8String(){
	s:="雨.痕"
	println(len(s),utf8.RuneCountInString(s))
}//7 3

func defineAArray(){
	var a[4] int
	b:=[4]int{2,5}
	c:=[4]int{5,3:10}//
	d:=[...] int{1,2,3}//according to the quantity of parameter
	e:=[...] int{10,3:100}
	fmt.Println(a,b,c,d,e)
}//[0 0 0 0] [2 5 0 0] [5 0 0 10] [1 2 3] [10 0 0 100]

type user struct{
	name string
	age byte
}
func omitArrayType(){
	d:=[...] user{
		{"Tom",20},
		{"Mary",18},
	}
	fmt.Printf("%#v\n",d)
}


func main() {
	//stringMain()
	//sliceString()
	//goThroughString()
	//appendJoin()
	//fmt.Print(appendJoinString())
	//testUnicode()
	//stringConvert()
	//lenUtf8String()
	//defineAArray()

	omitArrayType()



}
