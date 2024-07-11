package main

import (
	"log"
	"time"
)

const tickRate = 2 * time.Second

func main() {
	log.Println("start")
	ticker := time.NewTicker(tickRate).C // periodic
	stopper := time.After(5 * tickRate)  // one shot
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}
	log.Println("finish")
}
