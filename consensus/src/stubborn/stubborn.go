package stubborn

import "net"

// Interface to abstract the channel
type stubChannel interface {
	sSend(int, []byte)	bool
	sReceive()    		[]byte
}

func newStubChannel(peers map[int] *net.UDPAddr) (channel stubChannel){
	channel = newChannel(peers)
	return 
} 