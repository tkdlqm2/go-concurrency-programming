package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	
	var wg sync.WaitGroup


	// 여기서는 1을 인자로 Add를 호출해 하나의 고루틴이 시작된다는 것을 나타낸다.
	wg.Add(1)

	go func() {

		// 여기에서는 고루틴의 클로저를 종료하기 전에 waitGroup에게 종료한다고 알려주기 위해, defer키워드를 사용해 Done을 호출한다.
		defer wg.Done()
		fmt.Println("1st goroutine sleeping")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping")
		time.Sleep(2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("3nd goroutine sleeping")
		time.Sleep(3)
	}()


	// 여기서 wait 호출하는데, 이 호출로 인해 main 고루틴은 다른 모든 고루틴이 자신들이 종료되어다고 알릴 때 까지 대기한다.
	wg.Wait()
	fmt.Println("All goroutine complete")
	

}