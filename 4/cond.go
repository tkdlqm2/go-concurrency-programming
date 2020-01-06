package main

import (
	"fmt"
	"sync"
)

func main() {

	// ㅇㅕ기서 새로운 cond를 초기화 한다. NewCond 함수는 sync.Locker 인터페이스를 만족하는 타입을 인수로 받는다. 이 때문에
	// cond 타입은 동시에 실행해도 안전한 방식으로 손쉽게 다른 고루틴들과의 조정이 가능하다.
	c := sync.NewCond(&sync.Mutex{})

	// 여기서는 Locker를 이 상태로 고정시킨다. wait이 호출되면, wait호출에 진입할 때 자동적으로 Locker의 Unlock을 호출하기 때문에 이 작업이 필요하다.
	c.L.Lock()

	for conditionTrue() == false {

		// 여기서는 조건이 충족됬다는 알립을 기다린다. 이것은 대기하는 호출로서, 해당 고루틴은 일시 중지된다.
		c.Wait()
	}

	// 여기서는 이 조건에 대한 Locker의 잠금을 해제한다. wait 호출을 빠져나오면서 이 조건에 대한 Locker의 Lock을 호출하기 때문에 이 작업이 필요함.
	c.L.Unlock()
}