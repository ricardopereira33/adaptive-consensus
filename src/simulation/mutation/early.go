package mutation

import (
	con "simulation/consensus"
	stb "simulation/stubborn"
)

// Early is a mutation type
type Early struct {
	peer *con.Peer
}

// NewEarly creates a new early mutation
func NewEarly(peer *con.Peer) (early *Early) {
	early = new(Early)
	early.peer = peer

	return
}

// Delta0 is the delta0 implementation
func (early *Early) Delta0(id int, pack *stb.Package) bool {
	channel := early.peer.GetChannel()
	isFresh := fresh(channel.GetPackage(id), pack)
	isMajority := majority(pack, early.peer.GetNumberParticipants())

	needAck := early.peer.NeedAck(id)

	return (isFresh || isMajority) || needAck
}

// Delta is the delta implementation
func (early *Early) Delta(id int) bool {
	return true
}
