package main

import (

	"bytes"
	"fmt"
	"os"
)

func main() {
	

	//  출력이 얼마나 발생할지 모르기 때문에 여기서 메모리 내부에 버퍼를 생성한다.
	// 언제나 그런 것은 아니지만, Stdout에 직접 쓰는 것보다 조금 빠르다.
	var stdoutBuff bytes.Buffer

	// 프로세스가 종료되기 전에 버퍼가 stdout에 쓰여지도록 한다.
	defer stdoutBuff.WriteTo(os.Stdout)
	
	// 용량이 4인 버퍼 채널을 생성한다.
	intStream := make(chan int,4)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 4; i ++ {
			fmt.Fprintf(&stdoutBuff,"Sending: %d \n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v \n", integer)
	}
}