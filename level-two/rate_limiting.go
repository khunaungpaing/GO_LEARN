package leveltwo

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func usingLib(r rate.Limit, b int) {

	limiter := rate.NewLimiter(r, b)

	for i := 1; i <= 5; i++ {
		// Wait until the limiter allows the operation
		if err := limiter.Wait(context.Background()); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Perform the operation
		fmt.Printf("Operation %d performed at %s\n", i, time.Now().Format("15:04:05"))
	}
}

func notUsingLib(ctx context.Context, rate int, burstLimit int) {

	rateLimit := time.Second / time.Duration(rate)

	throttle := make(chan time.Time, burstLimit)

	go func() {
		ticker := time.NewTicker(rateLimit)
		defer ticker.Stop()
		for {
			select {
			case throttle <- time.Now():
			case <-ctx.Done():
				return
			}
			<-ticker.C // wait for the next tick
		}
	}()

	for i := 1; i <= 5; i++ {
		// Wait until the limiter allows the operation
		select {
		case <-throttle:
			// Perform the operation
			fmt.Printf("Operation %d performed at %s\n", i, time.Now().Format("15:04:05"))
		case <-ctx.Done():
			return // exit the loop when the context is canceled
		}
	}
}

func RateLimiterInGo() {
	fmt.Println("using lib")
	usingLib(5, 1)

	ctx := context.Background()
	rateLimit := 5
	burstLimit := 5

	fmt.Println("\nnot using lib")
	notUsingLib(ctx, rateLimit, burstLimit)
}
