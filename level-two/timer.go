package leveltwo

import (
	"fmt"
	"sync"
	"time"
)

func TimerInGo() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Create a timer that fires after 2 seconds
		timer := time.NewTimer(2 * time.Second)

		// The <-timer.C channel will block until the timer fires
		<-timer.C
		fmt.Println("Timer1 expired!")
	}()

	go func() {
		defer wg.Done()
		timer2 := time.NewTimer(2 * time.Second)

		// Stop the timer before it fires
		if !timer2.Stop() {
			// Drain the timer's channel if it has fired
			<-timer2.C
		}
		fmt.Println("Timer2 stopped!")
	}()

	wg.Wait()
}
