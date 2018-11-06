package mutation

import stb "simulation/stubborn"

// Ring is a mutation of a ring
type Ring struct { 
	channel stb.StubChannel
}

// NewRing creates a new ring
func NewRing(channel stb.StubChannel) (r *Ring) {
	r 		  = new(Ring)
	r.channel = channel

	return r
}

// Delta0 is the delta0 implementation
func (r Ring) Delta0(id int, pack *stb.Package) bool {
	isFresh    := fresh(r.channel.GetPackage(id), pack)
	isMajority := majority(pack, r.channel.GetNParticipants())
	isIDEqual  := id == ((r.channel.GetPeerID() % r.channel.GetNParticipants()) + 1)
	
	return isIDEqual && (isFresh || isMajority)
}
    
// Delta is the delta implementation
func (r Ring) Delta(id int) bool {
	return id == ((r.channel.GetPeerID() % r.channel.GetNParticipants()) + 1)
}