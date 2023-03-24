// golang
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	j := 0
	for i := 0; i < 990999099; i++ {
		j += i
	}
	fmt.Println(j)

	cost := time.Since(start)
	fmt.Printf("elapsed time：%s", cost)
}
