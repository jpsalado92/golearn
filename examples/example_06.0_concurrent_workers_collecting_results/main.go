package main

import (
	"fmt"
	"time"
)

type payload struct {
	text string
}

func requestPayloadGenerator() payload {
	return payload{text: "a"}
}

func payloadProcessor(p payload) {
	fmt.Println(p.text)
	time.Sleep(4 * time.Second)
	fmt.Println("done")
}

func main() {

	// payloadChan := make(chan payload, 4)

	pls := make([]payload, 0)
	for i := 0; i < 5; i++ {
		p := requestPayloadGenerator()
		pls = append(pls, p)
	}
	// fmt.Println(pls)
	for i := range pls {
		// fmt.Println(pls[i])
		go payloadProcessor(pls[i])
	}
	time.Sleep(5 * time.Second)
}
