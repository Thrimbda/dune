package main

import (
	"math"
	"syscall"
	"time"
)

const (
	SAMPLING_COUNT  = 200
	TOTAL_AMPLITUDE = 300
)

func main() {
	var busySpan [SAMPLING_COUNT]uint32
	amplitude := TOTAL_AMPLITUDE / 2
	radian := 0.0
	radianIncrement := 2.0 / SAMPLING_COUNT
	for i := 0; i < SAMPLING_COUNT; i++ {
		busySpan[i] = uint32((float64(amplitude) + (math.Sin(math.Pi*radian) * float64(amplitude))))
		radian += radianIncrement
	}
	var kernel = syscall.NewLazyDLL("Kernel32.dll")
	GetTickCount := kernel.NewProc("GetTickCount")

	var startTime uint32
	for j := 0; ; j = (j + 1) % SAMPLING_COUNT {
		r, _, _ := GetTickCount.Call()
		startTime = uint32(r)
		for uint32(r)-startTime <= busySpan[j] {
			r, _, _ = GetTickCount.Call()
		}
		time.Sleep(time.Duration(TOTAL_AMPLITUDE - busySpan[j]))
	}
}
