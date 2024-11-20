package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	numChannel := make(chan int, 10)
	stringChannel := make(chan string, 10)

	var wg sync.WaitGroup

	for val := 1; val <= 10; val++ {
		numChannel <- val
	}
	close(numChannel)

	for worker := 0; worker < 10; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			number := <-numChannel
			text := strconv.Itoa(number)
			stringChannel <- text
		}()
	}

	go func() {
		wg.Wait()
		close(stringChannel)
	}()

	fmt.Println("Результаты обработки:")
	for convertedText := range stringChannel {
		fmt.Println("Преобразовано:", convertedText)
	}
}
