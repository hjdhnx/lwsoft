// +build windows
//运行时请带上参数GOARCH=386

package mywin32api

import (
	"math/rand"
	"strconv"
	"time"
)

func Leftlick() {
	println("test")
}
func SleepStrMs(str_time string) {
	int_time, _ := strconv.Atoi(str_time)
	int_time /= 1000
	time.Sleep(time.Duration(int_time) * time.Second)
}

func SleepIntS(int_second int) {
	time.Sleep(time.Duration(int_second) * time.Second)
}
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
