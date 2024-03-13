package main

import(
	"fmt"
	"time"
)

type RateLimiter struct(){
	maxReqPerSec int
	currReqCount int
}

func NewRateLimiter(maxReqPerSec int) *RateLimiter{
	return &RateLimiter{
		maxReqPerSec: maxReqPerSec,
		currReqCount: 0,
	}
}

func (rl *RateLimiter) isAllowed() bool{
	if rl.maxReqPerSec > rl.currReqCount{
		rl.currReqCount++
		return true
	}
	return false
}

func main(){
	rl := NewRateLimiter(100)
	
	fmt.Printf("Hello world\n")
}