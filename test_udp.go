package main

import (
	"air-trail-backend/utils"
	"fmt"
	"log"
	"time"
)

func UDP_Mcast() {
	reader := utils.UdpMulticast{
		Group: "239.0.0.0",
		Port:  50000,
	}

	if err := reader.Join(); err != nil {
		log.Fatal("failed to join multicast:", err)
	}
	defer reader.Close()

	dataChan := make(chan []byte)

	go reader.ReadLoop(dataChan)

	for {
		select {
		case msg, ok := <-dataChan:
			if !ok {
				log.Println("channel closed, stopping")
				return
			}
			fmt.Printf("Received: %s\n", string(msg))
		case <-time.After(5 * time.Second):
			fmt.Println("no data in last 5s...")
		}
	}
}
