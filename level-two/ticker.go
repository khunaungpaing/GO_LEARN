package leveltwo

import (
	"fmt"
	"time"
)

func TickerInGo() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker : ", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Process Done!")
}
