package mutation

import (
	con "simulation/consensus"
	stb "simulation/stubborn"
)

// OldRing is a mutation of a ring
type OldRing struct {
	peer *con.Peer
}

// NewOldRing creates a new ring
func NewOldRing(peer *con.Peer) (ring *OldRing) {
	ring = new(OldRing)
	ring.peer = peer

	return
}

// Delta0 is the delta0 implementation
func (ring *OldRing) Delta0(id int, pack *stb.Package) bool {
	channel := ring.peer.GetChannel()
	isFresh := fresh(channel.GetPackage(id), pack)
	isMajority := majority(pack, ring.peer.GetNumberParticipants())
	isIDEqual := id == ((ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()) + 1)

	return isIDEqual && (isFresh || isMajority)
}

// Delta is the delta implementation
func (ring *OldRing) Delta(id int) bool {
	return id == ((ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()) + 1)
}
