package main

import(
	"fmt"
	"sync"

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

	var wg sync.WaitGroup
	for _, salutation := range [] string{"hello","greetings","good day"} {
		wg.Add(1)

		// 여기서 다른 함수처럼 매개 변수를 선언한다. 어떤 일이 발생하는지 보다 확실하게 하기 위해 원래의 salutaion 변수를 가리킨다.
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)

			// 여기서는 현재 반복 회차의 변수를 클로저로 전달한다. 문자열 구조체의 복사본이 만들어지므로, 고루틴이 실행될 때 적절한 문자열을 참조하리라는 것을 보장할 수 있다.
		}(salutation)
	}
	wg.Wait()


}