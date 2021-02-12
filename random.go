package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr, flag := []int{}, false
	arr = append(arr, rand.Intn(100))
	for i := 0; i < 180; i++ {
		n := rand.Intn(180)
		for j := 0; j < len(arr); j++ {
			if n == arr[j] {
				flag = true
			}
		}
		if !flag {
			arr = append(arr, n)
		} else {
			flag = false
		}
	}
	fmt.Println(arr[0:100])
	fmt.Println("Len:", len(arr))
}
