package main

// Chan can be used for Unidirectional data - Sending or Recieving  (<-)

func main() {
	var dataStream chan interface{} // Empty interface called as 'off type' interface
	dataStream = make(chan interface{})

	writeStream := make(chan<- interface{})
	readStream := make(<-chan interface{})

	<-writeStream
	readStream <- struct{}{}
}
