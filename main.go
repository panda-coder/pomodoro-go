package main

import (
	"fmt"
	"math"
	"time"

	"github.com/schollz/progressbar/v3"
)

const TOTAL_TIME = 100
const MINUTES = 25
const MINUTE_FACTOR = 60

func main() {
	bar := progressbar.Default(TOTAL_TIME)

	startTime := time.Now()
	finishTime := startTime.Add(time.Second * (MINUTE_FACTOR * MINUTES))

	for {
		now := time.Now()
		diff := finishTime.Sub(now)
		percent := (1 - (diff.Minutes() / MINUTES)) * 100
		bar.Set(int(math.Round(percent)))
		time.Sleep(10 * time.Second)
		if percent >= 100 {
			break
		}
	}
	fmt.Println("Executed")
}
