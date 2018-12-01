package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	fmt.Println("Start...")
	wg.Add(1)
	for i := 0; i < 4; i++ {
		k := i
		go func() {
			defer wg.Done()
			task(k)
		}()
	}

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Done...")
}

func task(num int) {
	fmt.Printf("Starting task %d\n", num)
	time.Sleep(time.Second * 1)
	fmt.Printf("Finished task %d\n", num)
}
