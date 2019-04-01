package mutation

import (
    stb "simulation/stubborn"
    con "simulation/consensus"
)

// Ring is a mutation of a ring
type Ring struct {
    peer *con.Peer
}

// NewRing creates a new ring
func NewRing(peer *con.Peer) (ring *Ring) {
    ring = new(Ring)
    ring.peer = peer

    return
}

// Delta0 is the delta0 implementation
func (ring Ring) Delta0(id int, pack *stb.Package) bool {
    channel := ring.peer.GetChannel()
    isFresh := fresh(channel.GetPackage(id), pack)
    isMajority := majority(pack, ring.peer.GetNumberParticipants())
    isIDEqual := id == ((ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()) + 1)

    return isIDEqual && (isFresh || isMajority)
}

// Delta is the delta implementation
func (ring Ring) Delta(id int) bool {
    return id == ((ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()) + 1)
}
