package mutation

import "consensus/stubborn"

// Early is a mutation type
type Early struct { 
    channel stubborn.StubChannel
}

// NewEarly creates a new early mutation
func NewEarly(channel stubborn.StubChannel) (e *Early) {
    e         = new(Early)
    e.channel = channel
    
    return e
}

// Delta0 is the delta0 implementation
func (e Early) Delta0(id int, pack *stubborn.Package) bool {
    isFresh    := fresh(e.channel.GetPackage(id), pack)
    isMajority := majority(pack, e.channel.GetNParticipants())
    
    return isFresh || isMajority
}
    
// Delta is the delta implementation
func (e Early) Delta(id int) bool {
    return true
}