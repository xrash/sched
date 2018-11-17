# sched

Small task scheduler written in Go.

## Usage example

```go
package main

import (
	"fmt"
	"github.com/xrash/sched"
	"time"
)

func test(prev, next time.Time) {
	fmt.Println("===")
	fmt.Println("prev", prev)
	fmt.Println("next", next)
}

func main() {
	loopScheduler := sched.NewLoopScheduler(time.Second * 2)

	j := sched.LaunchJob(test, loopScheduler)

	time.Sleep(time.Second * 10)

	j.Stop()
	fmt.Println("stopped")

	time.Sleep(time.Second * 5)
}
```
