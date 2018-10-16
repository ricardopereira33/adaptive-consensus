package stubborn

import (
	"time"
	"log"
	"net"
)

const ( 
	// MaxDatagramSize is the maximum size of a datagram packet
	MaxDatagramSize = 2048  
	// Debug flag, when it is true, prints some debug infos	
	Debug 	        = true	
	// MaxTries is the maximum value of tries
	MaxTries 		= 3
	// DefaultDelta is the default time to relay the messages to the others peers
	DefaultDelta	= time.Second * 3
)

// Channel to send and receive messages between peers
type Channel struct {
	peers 		map[int] *net.UDPAddr
	connection	*net.UDPConn
	outBuffer	Buffer
	inBuffer	chan *Package
	delta0Func	func(int, *Package) bool
	deltaFunc	func(int) bool
}

func newChannel(id int, ownPort string, neighborsPorts []string) (channel *Channel) {
	channel = new(Channel)

	channel.initConnection(id, ownPort)
	channel.peers	  = listOfPeers(neighborsPorts)
	channel.outBuffer = newBuffer(len(neighborsPorts))
	channel.inBuffer  = make(chan *Package)
	return
}

func (c Channel) initConnection(id int, port string) {
	if Debug { log.Println("Init Peer - 127.0.0.1:" + port) }

	addr, err := net.ResolveUDPAddr("udp", ":" + port)
    checkError(err, false)

    conn, err := net.ListenUDP("udp", addr)
    checkError(err, false)
    
	conn.SetReadBuffer(MaxDatagramSize)
	c.connection = conn
}

func (c Channel) delta0(id int, pack *Package) {
	c.delta0Func(id, pack)
}

func (c Channel) delta(id int) {
	c.deltaFunc(id)
}

func (c Channel) retransmission() {
	tries := 0
	for {
		time.Sleep(DefaultDelta)	
		for id := range c.peers {
			if c.deltaFunc(id) || tries > MaxTries {
				pack := c.outBuffer.getElem(id)
				
				if pack != nil && !pack.arrived { 
					c.send(id) 
				}
			}
			tries++
		}
	}
}


// Exported methods

// SetDelta0 is the method to define the delta0 implemention
func (c Channel) SetDelta0(f func(int, *Package) bool) {
	c.delta0Func = f
}

// SetDelta is the method to define the delta0 implemention
func (c Channel) SetDelta(f func(int) bool) {
	c.deltaFunc = f	
}

// Init is the method that start receipt of the message 
func (c Channel) Init() {
	go c.receive()
	go c.retransmission()
}

// Close is the method that closes the UDP connection 
func (c Channel) Close() {
	c.connection.Close()
}

// Auxiliary Functions 

func listOfPeers(ports []string) (list map[int] *net.UDPAddr){
	list = make(map[int] *net.UDPAddr)
	
	for id, port := range ports {
		addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:" + port)
		checkError(err, false)

		list[id] = addr
	}
	return
}