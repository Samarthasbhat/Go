package main


import (
		"fmt"
		"io/ioutil"
		"net/http"
		"sync"
		"time"
)


func fetchURL(url string, wg *sync.WaitGroup, ch chan<-string){
	defer wg.Done()

	start := time.Now() // time 
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return 
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	duration := time.Since(start)
	ms := int(duration.Milliseconds())
	if err != nil{
		ch <- fmt.Sprintf("Error reading response from %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("Fetched %s in %v (%d bytes)", url, ms , len(body))
}

func main(){
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(urls))


	for _, url := range urls{
		wg.Add(1)
		go fetchURL(url, &wg, ch)
	}

	wg.Wait()
	close(ch)

	for response := range ch{
		fmt.Println(response)
	}
}