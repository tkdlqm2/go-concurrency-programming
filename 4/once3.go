package main

import(

	"sync"
)

func main() {

	var onceA, onceB sync.Once

	var initB func()
	initA := func() {
		onceB.Do(initB)
	}

	// 이것이 호출될 때 2번의 호출이 리턴될 때 까지 진행되지 않는다.
	initB = func() {
		onceA.Do(initA) // 1 번
	}

	// DO 호출이 종료될 때까지 1번의 Do 호출이 진행되지 않기 때문에, 이 프로그램은 데드락에 빠지며 이는 데드락의 전형적인 예라고 할 수 있다.
	onceA.Do(initA) // 2 번
}