package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-ch)

	
}

func counter(ch chan<- int){
	for i:=0; ; i++ {
		ch <- i
	}
}


func main() {
	var nextID nextCh = make(chan int)



	go counter(nextID) // Start the counter in a goroutine
	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}