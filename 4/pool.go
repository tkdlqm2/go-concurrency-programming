package main

import (

	"fmt"
	"sync"
)

func main() {

	myPool := &sync.Pool {
		New: func() interface {} {
			fmt.Println("Creating new instance")
			return struct {}{}
		},
	}

	// 여기서는 풀의 Get을 호출한다. 인스턴스가 아직 초기화되지 않았기 때문에 이 호출은 풀에 정의된 New 함수를 호출한다.
	myPool.Get()
	instance := myPool.Get()

	// 여기서는 이전에 조회했던 인스턴스를 다시 풀에 돌려놓는다. 이는 사용 가능한 인스턴스의 수를 1로 증가시킨다.
	myPool.Put(instance)

	// 이 호출이 실행되면 이전에 할당됐다가 다시 풀에 넣은 인스턴스를 다시 사용한다. New 함수는 호출되지 않는다.
	myPool.Get()
}

