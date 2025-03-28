package main

import (
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
	"bufio"
)

func main(){
	for _, fname := range os.Args[1:]{
	
		var lc, wc, cc int
	
		file, err := os.Open(fname)

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			continue
		}


		scan := bufio.NewScanner(file)

		for scan.Scan(){
			s := scan.Text()

			wc += len(strings.Fields(s))  // additional number of words
			cc += len(s)
			lc++
		}

		//  data, err := ioutil.ReadAll(file) 
		 

		//  if err != nil{
		// 	fmt.Fprintln(os.Stderr, err)
		// 	continue
		// }
		// fmt.Println("The file has", len(data), " bytes")
		
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		
		file.Close()
	}
}