package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	// 每次都会打印相同数据
	fmt.Println("The secret number is ", secretNumber)
}
