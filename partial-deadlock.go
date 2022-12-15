package main

func prod(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func cons(ch1, ch2 chan int) {
	for {
		select {
		case x := <-ch1:
			print(x)
		case x := <-ch2:
			print(x)
		}
	}
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go prod(ch1)
	go prod(ch2)
	cons(ch1, ch1)
}
