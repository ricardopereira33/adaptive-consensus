package stubborn

// Interface to abstract the channel
type stubChannel interface {
	sReceive() *Package
	sSend(int, []byte)					  
	setDelta0(f func(int, *Package) bool)	
	setDelta(f func(int) bool)				
}

func newStubChannel(id int, ownPort string, neighborsPorts []string) (channel stubChannel){
	channel = newChannel(id, ownPort, neighborsPorts)
	return 
}
