package main

import "fmt"

func main() {
	if err := calc(); err != nil {
		fmt.Printf("Ошибка при выполнении: %s", err)
	}
}