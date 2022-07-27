package channel

import (
	"fmt"
	"sync"
	"time"
)

func ChannelDemo() {
	//Channel for goroutines communication
	ch := make(chan int, 5)

	//Waitgroup for coordination
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Sending data..")
		SendData(ch)
		fmt.Println("All data sent")
		wg.Done()
	}()

	go func() {
		fmt.Println("Waiting for some time before receiving..")
		time.Sleep(5 * time.Second)
		ReceiveData(ch)
		wg.Done()
	}()
	wg.Wait()
}

func SendData(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Data Sent:", i)
		ch <- i
	}
	close(ch)
}

func ReceiveData(ch <-chan int) {
	for msg := range ch {
		fmt.Println("Data received from channel:", msg)
		// if msg, ok := <-ch; ok {
		// 	fmt.Println("Data received from channel:", msg)
		// }
	}
	//close(ch)
}
