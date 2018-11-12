package stubborn

import (
    "sort"
    "time"
    "log"
    "net"
    ex "consensus/exception"
)

const ( 
    // MaxDatagramSize is the maximum size of a datagram packet
    MaxDatagramSize = 2048  
)	

var ( // Default value (adaptive) 
    // MaxTries is the maximum value of tries 
    MaxTries     = 3
    // DefaultDelta is the default time to relay the messages to the others peers 
    DefaultDelta = time.Second * 3
    // Debug flag, when it is true, prints some debug infos	
    Debug        = true	
)

// Channel to send and receive messages between peers
type Channel struct {
    Peers       map[int] *net.UDPAddr
    Connection	*net.UDPConn
    OutBuffer	Buffer
    InBuffer	chan *Package
    Delta0Func	func(int, *Package) bool
    DeltaFunc	func(int) bool
    PeerID		int
    CoordID		int
}

func newChannel(ownPort string, allPorts []string) (channel *Channel) {
    channel            = new(Channel)
    channel.Connection = initConnection(ownPort)
    channel.OutBuffer  = newBuffer(len(allPorts))
    channel.InBuffer   = make(chan *Package)
    channel.Peers, channel.PeerID = listOfPeers(ownPort, allPorts)
    
    return
}

func (c *Channel) delta0(id int, pack *Package) bool {
    return c.Delta0Func(id, pack)
}

func (c *Channel) delta(id int) bool {
    return c.DeltaFunc(id)
}

func (c *Channel) retransmission() {
    if Debug { 
        log.Println("Retransmisson start...")
    }

    tries := 0
    for {
        time.Sleep(DefaultDelta)	
        for id := range c.Peers {
            if c.delta(id) || tries > MaxTries {
                pack := c.OutBuffer.getElem(id)
                
                if pack != nil && !pack.Arrived { 
                    c.send(id) 
                }
            }
            tries++
        }
    }
}

/*** Exported methods ***/

// GetPeerID returns the peer ID
func (c *Channel) GetPeerID() int {
    return c.PeerID
}

// GetPackage returns the last package sent to id
func (c *Channel) GetPackage(id int) *Package {
    pack := c.OutBuffer.getElem(id)

    return pack
}

// GetNParticipants returns the number of participants
func (c *Channel) GetNParticipants() int {
    return len(c.Peers)
}

// GetCoordID returns the coordinator ID
func (c *Channel) GetCoordID() int {
    return c.CoordID
}

// SetDelta0 is the method to define the delta0 implemention
func (c *Channel) SetDelta0(f func(int, *Package) bool) {
    c.Delta0Func = f
}

// SetDelta is the method to define the delta0 implemention
func (c *Channel) SetDelta(f func(int) bool) {
    c.DeltaFunc = f	
}

// SetCoordinator saves the coordainator ID
func (c *Channel) SetCoordinator(coordID int) {
    c.CoordID = coordID
}

// SetMaxTries sets the MaxTries value.
func (c *Channel) SetMaxTries(max int) {
    MaxTries = max
}

// SetDefaultDelta sets the DefaultDelta value.
func (c *Channel) SetDefaultDelta(ddelta int) {
    DefaultDelta = time.Second * time.Duration(ddelta)
}

// Init is the method that start receipt of the message 
func (c *Channel) Init() {
    go c.receive()
    go c.retransmission()
}

// Close is the method that closes the UDP connection 
func (c *Channel) Close() {
    c.Connection.Close()
}

func (c Channel) printStatus() {
    log.Println(c.Peers) 	
    log.Println(c.Connection) 	
    log.Println(c.OutBuffer) 	
    log.Println(c.Delta0Func != nil) 	
    log.Println(c.DeltaFunc != nil) 	
}

// Auxiliary Functions 

func initConnection(port string) (conn *net.UDPConn){
    if Debug { 
        log.Println("Init Peer - 127.0.0.1:" + port) 
    }

    addr, err := net.ResolveUDPAddr("udp", ":" + port)
    conn, err  = net.ListenUDP("udp", addr)
    ex.CheckError(err)
    
    conn.SetReadBuffer(MaxDatagramSize)

    return
}

func listOfPeers(ownPort string, ports []string) (list map[int] *net.UDPAddr, ownID int){
    list = make(map[int] *net.UDPAddr)
    sort.Strings(ports)

    for index, port := range ports {
        if port ==  ownPort {
            ownID = index+1
        } else {
            addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:" + port)
            ex.CheckError(err)
            
            list[index+1] = addr
        }
    }
    return
}