package main

import "fmt"

func main() {

	arrayChannels := make([]chan int, 3)
	for i := 0; i < len(arrayChannels); i++ {
		arrayChannels[i] = make(chan int, 10)
	}

	for _, channel := range arrayChannels {
		go send(channel)
	}

	for _, channel := range arrayChannels {
		recv(channel)
	}

}

func send(a chan int){
	for val := 0; val < 3; val++ {
		a <- val
	}
	close(a)
}

func recv(a chan int) {
	for v := range a {
		fmt.Println(v)
	}
}
