package main 


import (
	"fmt"
	"os"
	"bytes"
)


func main() {
	var stdoutBuff bytes.Buffer  //Here we create in memory buffer to help the mitigate the non-determinastic nature of the output. It doesn't give any guarantees, but little faster than writing the stdout directly.
	defer stdoutBuff.WriteTo(os.Stdout) // Here we ensure the buffer is written out to stdout before the process exits

	intStream := make(chan int, 4) // created buffer channel with a capacity of one.

	go func ()  {
	
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i:= 0; i< 4; i++{
			fmt.Fprintf(&stdoutBuff," sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Recieved: %v\n", integer)
	}
}