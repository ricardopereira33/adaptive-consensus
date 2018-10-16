package main

import (
	"fmt"
	"stubborn"
	"encoding/json"
)

func handleMessages(channel stubborn.StubChannel) {

	for {
		message := channel.SReceive()
		
		message.PrintPacket()

		fmt.Println("--------")
		data := message.GetData()

		var msg Message 
		err := json.Unmarshal(data, &msg)
		checkError(err, false)

		fmt.Print("[ " + msg.Username + " ]")
		fmt.Println(msg.Info)
	}
}