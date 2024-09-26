package main

import (
	"log"
)

func main() {
	if err := calc(); err != nil {
		log.Fatal("Ошибка при выполнении: %s", err)
	}
}