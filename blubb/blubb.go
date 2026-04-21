package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/dustin/go-broadcast"
)

func main() {
	broadcaster := broadcast.NewBroadcaster(1)

	ch := make(chan interface{})
	broadcaster.Register(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("received msg, going to sleep for 5 seconds...")
				time.Sleep(5 * time.Second)

			case <-time.After(time.Second * 11):
				return
			}
		}
	}()

	fmt.Println(broadcaster.TrySubmit(0))
	time.Sleep(time.Second * 1)
	fmt.Println(broadcaster.TrySubmit(1))
	time.Sleep(time.Second * 1)
	fmt.Println(broadcaster.TrySubmit(2))
	time.Sleep(time.Second * 1)
	fmt.Println(broadcaster.TrySubmit(3))

	wg.Wait()
}
