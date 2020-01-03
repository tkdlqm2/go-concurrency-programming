package main

import (
	"fmt"
	"sync"
	"time"
	"sync/atomic"
	"bytes"
)

func main() {

	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// tryDir는 어떤 사람이 특정 방향으로 움직이게 시도할 수 있도록 해주고, 그 시도가 성공했는지 아닌지를 리턴하는 함수이다. 각 방향은 그 방향으로 움직이고자 하는 사람들의 수를 표시한다.
	tryDir := func(dirName string, dir * int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%v", dirName)

		// 먼저, 인자로 주어진 방향을 1씩 증가시킴으로써 그 방향으로 움직이겠다는 의사를 표현한다.
		atomic.AddInt32(dir,1)

		// takeStep은 모든 참가자가 일정한 보조로 움직이는 것을 시뮬레이션 된다.
		takeStep()
		if atomic.LoadInt32(dir) == 1{
			fmt.Fprintf(out, "Success!")
		}
		takeStep()

		// 이 시점에서 이 사람은 이 방향으로 갈 수 없다는 것을 깨닫고 포기한다. 그 방향을 1만큼 줄이는 것으로 이를 표현한다.
		atomic.AddInt32(dir,-1)
		return false
	}

	var left,right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left",&left,out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right",&right, out)
	}

	walk := func(walking * sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot : ", name)

		// 프로그램이 종료될 수 있도록 인위적인 제한을 두었다. 라이브락이 있는 프로그램은 이러한 제한이 없기 때문에 문제가 된다!
		for i := 0; i < 5; i++ {

			// 먼저 한 사람은 왼쪽으로 움직이려고 시도하고, 실패하면 오른쪽으로 움직이려 한다.
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}

		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	// 이 변수는 두 사람이 서로를 지나쳐 가거나 혹은 포기할 때 까지 프로그램이 기다릴 수 있게 하는 방법을 제공한다.
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway,"Alice")
	go walk(&peopleInHallway,"Barbara")
	peopleInHallway.Wait()
}