package main

import (
	"stubborn"
)

func handleMessages(channel stubborn.StubChannel) {
	for {
		pack 	:= channel.SReceive()
		message := bytesToMessage(pack)

		printMessage(message)
	}
}