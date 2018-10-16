package stubborn

// StubChannel is a interface to abstract the channel
type StubChannel interface {
	SReceive() *Package
	SSend(int, []byte)					  
	SetDelta0(f func(int, *Package) bool)	
	SetDelta(f func(int) bool)	
	Init()	
	Close()		
}

// NewStubChannel is the constructor of a stubborn channel
func NewStubChannel(id int, ownPort string, neighborsPorts []string) (channel StubChannel){
	channel = newChannel(id, ownPort, neighborsPorts)
	return 
}
