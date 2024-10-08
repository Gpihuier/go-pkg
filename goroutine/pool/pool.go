package pool

import "fmt"

func Channel(count int) {

	ch := make(chan struct{}, count)

	do := func() {
		fmt.Println("do")

		<-ch
	}

	go func() {	
		ch <- struct{}{}

		// do
		do()
	}()
}
