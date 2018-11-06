package mutation

import "consensus/stubborn"

type Early struct { 
	channel stubborn.StubChannel
}

func NewEarly(channel stubborn.StubChannel) (e *Early) {
	e 		  = new(Early)
	e.channel = channel
	
	return e
}

func (e Early) Delta0(id int, pack *stubborn.Package) bool {
	isFresh    := fresh(e.channel.GetPackage(id), pack)
	isMajority := majority(pack, e.channel.GetNParticipants())
	
	return isFresh || isMajority
}
	
func (e Early) Delta(id int) bool {
	return true
}