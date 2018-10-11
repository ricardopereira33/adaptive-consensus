package main

import (
    "log"
    "net"
    "encoding/json"
)

func reader(listener *net.UDPConn){
    if debug { log.Println("reader") }

    buffer := make([]byte, MaxDatagramSize)

    for {
        nBytes, srcAddr, err := listener.ReadFromUDP(buffer)

        if checkError(err, true){
            go handlePacket(buffer, nBytes, srcAddr)
        }
    }
}

//receives message from peer
func handlePacket(buffer []byte, nBytes int, srcAddr *net.UDPAddr){
    if debug { log.Println("handlePacket") }
    
    var msg Message
    
    data := buffer[:nBytes]
    err  := json.Unmarshal(data, &msg)
    checkError(err, false)

    printMessage(msg)
}