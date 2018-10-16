package main

import (
	"fmt"
	"stubborn"
	"encoding/json"
)

func handleMessages(channel stubborn.StubChannel) {

	for {
		message := channel.SReceive()
		
		
		var msg Message 
		data := message.GetData()
		err  := json.Unmarshal(data, &msg)
		checkError(err, false)

		fmt.Print("[ " + msg.Username + " ]")
		fmt.Println(msg.Info)
	}
}