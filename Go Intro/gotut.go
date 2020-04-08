package main

import (
	"fmt"
)

func foo() {
  defer wg.Done()
  for i := 0; i < 3; i++ {
    time.Sleep(100*time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}
