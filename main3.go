package main

import "fmt"

func main() {
	fmt.Println("Демонстрация блокировки канала:")
	blockingChan := make(chan int)
	fmt.Println("Ожидание чтения из канала...")
	data := <-blockingChan
	fmt.Println("Это не будет напечатано:", data)
}
