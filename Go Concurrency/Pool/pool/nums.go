package pool

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Nums() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated +=1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

for i:= numWorkers ; i>0;i--{
	go func ()  {
		defer wg.Done()
		mem := calcPool.Get().(*[]byte)
		defer calcPool.Put(mem)
	}()
}
wg.Wait()
fmt.Printf("%d calculators were created.", numCalcsCreated)
}


func connectToService() interface{} {
	time.Sleep(1*time.Second)
	return struct{}{}
}

//  If we started reuesting a new connection to the service. This is a network handler that opens connection to the service,

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func () {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil{
			log.Fatalf("Cannot listen %v", err)
		}

		defer server.Close()

		wg.Done()

		for{
			conn, err := server.Accept()
			if err != nil {
				log.Printf("Cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}