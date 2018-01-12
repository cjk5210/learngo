package main
type Integer int

func (a Integer)Less(b Integer)bool{
	return a<b
}

func (a *Integer)Add(b Integer){
	*a+=b
}

type LessAdder interface {
	Less(b Integer)bool
	Add(b Integer)
}
func main() {
	var a Integer=1
	println(a)
	var b LessAdder=&a

	println(b)

	//var c LessAdder=a  //enable this line will show error :cannot use a (type Integer) as type LessAdder in assignment:
														//Integer does not implement LessAdder (Add method has pointer receiver)
	println(c)
}
//It means if a interface contains a method used by pointer,you must use a struct's or type's pointer to give to the interface variable.
//Although the type Integer haven't a method like 'func (a *Integer)Less(b Integer)bool{ return *a<b }',it will generate a new method like:
/*
func (a *Integer) Less( b Integer) bool { return (*a). Less( b) }
we can get the value by a pointer,but we cann't get a pointer by a value
方法接收参数是通过值传递的，所以通过接收到的值是一个原值的copy，所以不能通过值获取原指针
如果接收的参数是一个指针，就可以通过指针来获取原值，并且可以改变原值

 */