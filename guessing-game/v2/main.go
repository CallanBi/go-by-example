package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	// 加上随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
