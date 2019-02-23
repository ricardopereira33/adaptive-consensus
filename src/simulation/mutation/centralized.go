package mutation

import (
    con "simulation/consensus"
    stb "simulation/stubborn"
)

// Centralized is a mutation type
type Centralized struct {
	peer *con.Peer
}

// NewCentralized creates a new centralized mutation
func NewCentralized(peer *con.Peer) (centralized *Centralized) {
	centralized = new(Centralized)
	centralized.peer = peer

	return
}

// Delta0 is the delta0 implementation
func (centralized Centralized) Delta0(id int, pack *stb.Package) bool {
	coordID := centralized.peer.GetCoordID()
	peerID := centralized.peer.GetPeerID()

    channel := centralized.peer.GetChannel()
	isCoord := id == coordID || peerID == coordID
	isFresh := fresh(channel.GetPackage(id), pack)
	isMajority := majority(pack, centralized.peer.GetNumberParticipants())

	return isCoord && (isFresh || isMajority)
}

// Delta is the delta implementation
func (centralized Centralized) Delta(id int) bool {
	return true
}
