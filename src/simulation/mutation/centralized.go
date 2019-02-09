package mutation

import "simulation/stubborn"

// Centralized is a mutation type
type Centralized struct {
	channel stubborn.StubChannel
}

// NewCentralized creates a new centralized mutation
func NewCentralized(channel stubborn.StubChannel) (centralized *Centralized) {
	centralized = new(Centralized)
	centralized.channel = channel

	return
}

// Delta0 is the delta0 implementation
func (centralized Centralized) Delta0(id int, pack *stubborn.Package) bool {
	coordID := centralized.channel.GetCoordID()
	peerID := centralized.channel.GetPeerID()

	isCoord := id == coordID || peerID == coordID
	isFresh := fresh(centralized.channel.GetPackage(id), pack)
	isMajority := majority(pack, centralized.channel.GetNumberParticipants())

	return isCoord && (isFresh || isMajority)
}

// Delta is the delta implementation
func (centralized Centralized) Delta(id int) bool {
	return true
}
