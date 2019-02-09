package mutation

import stb "simulation/stubborn"

// Ring is a mutation of a ring
type Ring struct {
	channel stb.StubChannel
}

// NewRing creates a new ring
func NewRing(channel stb.StubChannel) (ring *Ring) {
	ring = new(Ring)
	ring.channel = channel

	return 
}

// Delta0 is the delta0 implementation
func (ring Ring) Delta0(id int, pack *stb.Package) bool {
	isFresh := fresh(ring.channel.GetPackage(id), pack)
	isMajority := majority(pack, ring.channel.GetNumberParticipants())
	isIDEqual := id == ((ring.channel.GetPeerID() % ring.channel.GetNumberParticipants()) + 1)

	return isIDEqual && (isFresh || isMajority)
}

// Delta is the delta implementation
func (ring Ring) Delta(id int) bool {
	return id == ((ring.channel.GetPeerID() % ring.channel.GetNumberParticipants()) + 1)
}
