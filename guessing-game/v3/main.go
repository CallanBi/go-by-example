package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	// 读取一行
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// 去掉换行符
	input = strings.Trim(input, "\r\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
}
