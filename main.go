package main

import (
	"fmt"
	"time"

	"github.com/cshenton/oneten/oneten"
)

func main() {
	val := make([]bool, 200)
	n := 5000

	ons := []int{1, 2, 3, 4, 5, 6, 49, 53, 68, 111, 165, 167, 169, 199}
	for _, i := range ons {
		val[i] = true
	}

	c, _ := oneten.New(val, false, false)
	for i := 0; i < n; i++ {
		c.Next()
		fmt.Println(c)
		time.Sleep(10 * time.Millisecond)
	}
}
