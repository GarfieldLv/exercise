package main

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x) //receive data until channel is closed
	}
	done <- true // notify main done
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i // send data
	}

	close(data) //close channel
}

func main() {
	done := make(chan bool) // close sign to receive consumer
	data := make(chan int)  //data channel

	go consumer(data, done) //start consumer
	go producer(data)       // start

	<-done
}
