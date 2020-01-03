package main

import (
	"fmt"
	"sync"
	
)

func main() {

	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}

	wg.Add(1)
	go sayHello()

	// 합류지점
	wg.Wait()


}