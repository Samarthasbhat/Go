//  pool usually refers to a resource pool — a way to reuse expensive-to-create objects like database connections, goroutines, buffers, etc., to improve performance and reduce memory pressure.

package main

import (
	"fmt"
	"sync"
	"pool-example/pool"
)

func main() {
	pool.Nums()
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instances")
			return struct{}{}
		},
	}

	myPool.Get()             // 1. Get the Pool. It will invoke the new function defined on the pool since instances haven't it yet been instanstiated
	instance := myPool.Get() //1.
	myPool.Put(instance)     // 2. Here we put an instance previously retrieved back in the pool, This increase the available number of instances to one
	myPool.Get()             // 3. We will reuse the instances which were allocated previously and put it back in the pool. The new function will not be invoked

	
}

// Can we use any instantiative object? No, the go lang containes garbage collector  it will cleaned up instantly
