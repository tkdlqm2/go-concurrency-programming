package main

import (
	"fmt"
	"sync"
)

func main() {

	var numCalcsCreated int
	calcPool := &sync.Pool {
		New: func() interface{}{
			numCalcsCreated += 1
			mem := make([]byte, 1024)

			// byte 슬라이스들의 주소를 저장하고 있음에 유의하라.
			return &mem
		},
	}

	// 4KB로 풀을 시작한다.

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i -- {
		go func() {
			defer wg.Done()
			// 그리고 여기서는 타입이 당연히 byte 슬라이스에 대한 포인터라고 가정하고 있다.
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}