package stubborn

// StubChannel is a interface to abstract the channel
type StubChannel interface {
	SReceive() *Package
	SSend(int, []byte)
	SSendAll([]byte)					  
	SetDelta0(f func(int, *Package) bool)	
	SetDelta(f func(int) bool)
	SetMaxTries(int)
	SetDefaultDelta(int)	
	GetPeerID() int
	Init()	
	Close()		
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(ownPort string, allPorts []string) (channel StubChannel){
	channel = newChannel(ownPort, allPorts)
	return 
}
