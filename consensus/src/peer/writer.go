package main

import (
	"log"
	"bufio"
	"os"
	"encoding/json"
	"net"
)

func writer(conn *net.UDPConn) {
	if debug { log.Println("writer") }

	reader := bufio.NewReader(os.Stdin)
	var msg *Message

	for {
		input, err := reader.ReadString('\n')
		checkError(err, false)

		msg = createMessage(username, "Test-" + username + ": "+ input)
		send(conn, msg)
	}
}


// Sends message to all peers
func send(conn *net.UDPConn, msg *Message) {
	if debug { log.Println("send_by-" + username) }

	for _, peerAddr := range listConnections {
		jsonMsg, err  := json.Marshal(*msg)
		checkError(err, false)
		
		conn.WriteTo(jsonMsg, peerAddr)
	}
}