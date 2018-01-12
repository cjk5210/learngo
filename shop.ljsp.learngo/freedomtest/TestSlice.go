package main

import "fmt"

func main() {
	a:=[5]string{"a","b","c","d","e"}

	b:=a[:1]

	fmt.Printf("a--type:%T,%v \n",a,a)
	fmt.Printf("b--%v\n",b)
	c:=append(b,"f")

	fmt.Printf("a--type:%T,%v \n",a,a)
	fmt.Printf("b--%v\n",b)
	fmt.Printf("c--%v\n",c)
}
