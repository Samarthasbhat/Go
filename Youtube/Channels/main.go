package main

import (
	"log"
	"net/http"
	"time"
)

//  We can only close the channel once 

type result struct{
	url  string
 	err  error
	latency time.Duration
}

func get(url string, ch chan <-result){   // Write data to channel
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	}else{
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()   // Socket got closed
	}
}

func main() {
		results := make(chan result)
		list := []string{
			"https://www.google.com",
			"https://www.facebook.com",
			"https://www.youtube.com",
		}

		for _, url := range list {
			go get(url, results)   // go keyword with the "get" function
		}

		for range list {  // Need to read all results
			r :=  <-results // Read data from channel
			if r.err != nil {
				log.Printf("%-2s %s\n", r.url, r.err)
			}else{
				log.Printf("%-2s %s\n", r.url, r.latency)

			}
		}
}