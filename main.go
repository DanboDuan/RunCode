package main

import (
	"fmt"

	"github.com/leetcode/go/hello"
)

func main() {
	callTime, token := hello.GenerateToken()
	fmt.Println(callTime)
	fmt.Println(token)
}
