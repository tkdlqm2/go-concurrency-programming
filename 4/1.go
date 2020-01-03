package main

import(
	"fmt"
	"sync"
	"runtime"

)

func main() {
	// var wg sync.WaitGroup
	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcome"
	// }()
	// wg.Wait()
	// fmt.Println(salutation)

	// var wg sync.WaitGroup
	// for _, salutation := range []string{"hello","greetings","good day"} {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()

	// 		// 여기서는 문자열 슬라이스의 범위에 의해 생성된 반복문의 변수 salutation을 참조한다.
	// 		fmt.Println(salutation)
	// 	}()
	// }
	// wg.Wait()

	// var wg sync.WaitGroup
	// for _, salutation := range [] string{"hello","greetings","good day"} {
	// 	wg.Add(1)

	// 	// 여기서 다른 함수처럼 매개 변수를 선언한다. 어떤 일이 발생하는지 보다 확실하게 하기 위해 원래의 salutaion 변수를 가리킨다.
	// 	go func(salutation string) {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)

	// 		// 여기서는 현재 반복 회차의 변수를 클로저로 전달한다. 문자열 구조체의 복사본이 만들어지므로, 고루틴이 실행될 때 적절한 문자열을 참조하리라는 것을 보장할 수 있다.
	// 	}(salutation)
	// }
	// wg.Wait()

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <- chan interface{}
	var wg sync.WaitGroup

	// 측정을 위해서는 메모리상에 다수의 고루틴이 유지되야 하며, 이를 위해 고루틴이 절대로 종료되지 않도록 해야 한다.
	// 하지만 지금은 이 조건을 어떻게 수행할지 걱정하지 않아도 된다. 다만 이 고루틴이 프로세스가 끝날 때 까지 종료되지 않는다는 것만 알아두자.
	noop := func() {wg.Done(); <- c}


	// 여기서 생성할 고루틴의 수를 정의한다. 고루틴 하나의 크기에 점근적으로 접근하기 위해 대수의 범칙을 사용할 것이다.
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	
	// 여기서는 고루틴들을 생성하기 전에 소비된 메모리 양을 측정한다.
	for i := numGoroutines; i > 0; i -- {
		go noop()
	}
	wg.Wait()

	// 그리고 여기서 고루틴들을 생성한 후에 소비된 메모리의 양을 측정한다. 결과는 다음과 같다.
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after - before)/numGoroutines/1000)
}