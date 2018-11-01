package stubborn

// StubChannel is an interface to abstract the channel
type StubChannel interface {
	SReceive() *Package
	SSend(int, []byte)
	SSendAll([]byte)					  
	Init()	
	Close()	

	// Sets
	SetDelta0(f func(int, *Package) bool)	
	SetDelta(f func(int) bool)
	SetMaxTries(int)
	SetDefaultDelta(int)
	SetCoordinator(int)	
	
	// Gets
	GetPeerID() int
	GetCoordID() int
	GetPackage(int) *Package
	GetNParticipants() int 	
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(ownPort string, allPorts []string) (channel StubChannel){
	channel = newChannel(ownPort, allPorts)
	return 
}
