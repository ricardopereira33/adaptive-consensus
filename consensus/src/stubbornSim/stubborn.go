package stubborn

// StubChannelSim is an interface to abstract the channel
type StubChannelSim interface {
	SReceive() *Package
	SSend(int, []byte)
	SSendAll([]byte)					  
	Init()	

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

// NewStubChannelSim is the constructor of a stubborn channel
func NewStubChannelSim(peerID int) (channel StubChannelSim){
	channel = newChannelSim(peerID)
	return 
}
