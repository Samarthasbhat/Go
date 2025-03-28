package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	for _, fname := range os.Args[1:]{
		file, err := os.Open(fname)

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		 data, err := ioutil.ReadAll(file) 
		 
		 if err != nil{
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println("The file has", len(data), " bytes")
		file.Close()
	}
}