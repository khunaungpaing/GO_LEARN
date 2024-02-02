package leveltwo

import "fmt"

func ChannelBuffer() {
	bufch := make(chan int, 3)

	bufch <- 1
	bufch <- 2
	bufch <- 3
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case <-bufch:
			case v := <-bufch:
				fmt.Println(v)
			}

		}

	}()
}
