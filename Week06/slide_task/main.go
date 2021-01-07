package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i <10 ; i++ {

		fmt.Printf("随机数：%d\n",  rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10))
	}
}
