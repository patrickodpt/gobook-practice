package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for ind, arg := range os.Args[1:] {
		fmt.Println(arg)
		fmt.Println(ind)
	}
	mSecs := time.Since(start).Seconds()
	fmt.Printf("%g", mSecs)
}
