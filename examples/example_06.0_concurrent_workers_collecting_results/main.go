package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type payload struct {
	id  uuid.UUID
	num int
}

func requestPayloadGenerator(num int) payload {
	return payload{
		id:  uuid.New(),
		num: num,
	}
}

func payloadProcessor(p payload) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Processing task `%d` with id `{%s}`\n", p.num,p.id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task `%d` done\n", p.num)
}

func main() {

	// payloadChan := make(chan payload, 4)

	pls := make([]payload, 0)
	for i := 0; i < 5; i++ {
		p := requestPayloadGenerator(i)
		pls = append(pls, p)
	}
	// fmt.Println(pls)
	for i := range pls {
		// fmt.Println(pls[i])
		go payloadProcessor(pls[i])
	}
	time.Sleep(5 * time.Second)
}
