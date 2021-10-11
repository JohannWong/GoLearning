package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lesson13() {
	for i := 0; i < 5; i++ {
		rand_int := rand.Int()
		fmt.Println(rand_int)
	}
	for i := 0; i < 5; i++ {
		rand_intn := rand.Intn(8)
		fmt.Println(rand_intn)
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 5; i++ {
		fmt.Println(100 * rand.Float64())
	}
}
