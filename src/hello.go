package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("hello world: " + now.String())
}
