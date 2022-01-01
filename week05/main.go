package main

import (
	"rollingwin/rollingwindow"
	"sync"
	"time"
)

func main() {
	wins, err := rollingwindow.NewWindows(3*time.Second, time.Second)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		wins.WindowCountTimer()
	}()
	wins.AddRequest()
	time.Sleep(1 * time.Second)
	wins.AddRequest()
	time.Sleep(1 * time.Second)
	wins.AddRequest()
	wg.Wait()
}
