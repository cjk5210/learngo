package main


func consumer(data chan int,done chan bool){
	for x:=range data{
		println("recv:",x)
	}
	done<-true
}
func producer(data chan int){
	for i:=0;i<4;i++{
		data<-i			//send the num to channel
	}
	close(data)  //close channel
}
func main() {
	done:=make(chan bool)
	data:=make(chan int)
	go consumer(data,done)
	go producer(data)
	<-done		//block the main-process until the msg return from the channel
}
