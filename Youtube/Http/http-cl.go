package main

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"io/ioutil"
	"os"
)


// func handler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Hello , world! from %s\n", r.URL.Path[1:]) // Response writer
// }

const url = "https://jsonplaceholder.typicode.com"

type todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main(){
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))


	// Client
	resp, err := http.Get(url + "/todos/1")

	if err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}


	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		// fmt.Println(string(body))

		var item todo
		
		err = json.Unmarshal(body, &item)

		
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item)

	}
}