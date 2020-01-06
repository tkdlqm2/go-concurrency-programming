package main

import (
	"fmt"
	"sync"
)

func main() {

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 계속 진행해도 된다고 할 때 까지 고루틴은 여기서 대기
			<- begin
			fmt.Printf("%v has begun \n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines ...")

	close(begin)
	wg.Wait()

}