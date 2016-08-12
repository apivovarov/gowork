package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("max: ", max)
	t1 := time.Now().UnixNano()

	sieve := make([]bool, max)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0], sieve[1] = false, false

	lim := int(math.Sqrt(float64(max)))

	for n := 2; n < lim; n++ {
		if sieve[n] {
			for j := 2 * n; j < max; j += n {
				sieve[j] = false
			}
		}
	}

	t2 := time.Now().UnixNano()
	tt := float64(t2 - t1) / 1000000

	fmt.Printf("time: %.3f ms\n", tt)
}
