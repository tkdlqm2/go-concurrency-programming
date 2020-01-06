package main

import (

	"fmt"
	"sync"
	
)


// clicked이라는 조건을 가지고 있는 Button 타입을 정의한다.
type Button struct {
	Clicked *sync.Cond
}

func main() {
	
	button := Button{ Clicked: sync.NewCond(&sync.Mutex{})}

	// 여기서는 조건의 신호들을 처리하는 함수를 등록할 수 있는 편의 함수를 정의한다.
	// 각 핸들러는 자체 고루틴에서 실행되며, 고루틴이 실행 중이라는 것을 확인하기 전까지 subscribe는 종료되지 않는다.
	subscribe := func(c *sync.Cond, fn func()) {

		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	// 마우스 버튼이 올라왔을 때의 핸들러를 설정한다. 이는 결국 Clicked cond에서 Broadcast를 호출해 모든 핸들러에게
	// 마우스 버튼이 클릭됬음을 알린다. (보다 안정적인 구현을 위해서는 먼저 버튼이 눌렸는지 확인하면 된다.)
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	

	//여기서는 waitgroup을 생성한다. 이는 stdout에 대한 쓰기가 발생하기 전에 프로그램이 종료되지 않도록 하기 위해서이다.
	subscribe(button.Clicked, func(){
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})


	// 여기서는 버튼을 클릭할 때 버튼의 윈도우를 최대화하는 것을 시뮬레이션하는 핸들러를 등록한다.
	subscribe(button.Clicked, func(){
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})


	//여기서는 마우스가 클릭됬을 때 대화 상자를 표시하는 것을 시뮬레이션하는 핸들러를 등록한다.
	subscribe(button.Clicked, func(){
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})


	// 다음으로 사용자가 애플리케이션의 버튼을 클릭했다가 떼었을 때를 시뮬레이션 한다.
	button.Clicked.Broadcast()

	clickRegistered.Wait()

}