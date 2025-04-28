package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

)

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err!=nil{
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping: %d got %d\n", i,resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body
} 

func main() {
	 var (
		output io.WriteCloser = os.Stdout
		err error
		cnt int
		fails int
		data []byte
	 )

	 if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		defer output.Close()
	 }

	//  the output will be in the form of a JSON array,
	//  so add brackets before and after

	 fmt.Fprintln(output, "[")
	 defer fmt.Fprintln(output, "]")


	//  stop if we get two 404s in a row(get passed #404)

	for i:= 1; fails < 2 ; i++{
		if data = getOne(i); data == nil {
			fails++
			continue
		} 
	

	if cnt > 0 {
		fmt.Fprint(output, ",") // comma after the first one
	}

	_, err = io.Copy(output, bytes.NewBuffer(data))


	if err != nil {
		fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
		os.Exit(-1)
	}

	fails = 0
	cnt++
}

	 fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}