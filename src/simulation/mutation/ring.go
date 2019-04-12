package mutation

import (
	con "simulation/consensus"
    stb "simulation/stubborn"
)

// Ring is a mutation of a ring
type Ring struct {
    peer  *con.Peer
    steps int
}

// NewRing creates a new ring
func NewRing(peer *con.Peer) (ring *Ring) {
	ring = new(Ring)
    ring.peer = peer
    ring.steps = 0

	return
}

// Delta0 is the delta0 implementation
func (ring *Ring) Delta0(id int, pack *stb.Package) bool {
	channel := ring.peer.GetChannel()

    isFresh := fresh(channel.GetPackage(id), pack)
	isMajority := majority(pack, ring.peer.GetNumberParticipants())

    next := ring.next()
    previous := ring.previous()

	return (id == next || id == previous) && (isFresh || isMajority)
}

// Delta is the delta implementation
func (ring *Ring) Delta(id int) bool {
    lastPackage := ring.peer.GetChannel().LastPackageBuffered(id)

    if !lastPackage.Arrived {
        ring.steps++
    }

    next := ring.nextWithStep()
    previous := ring.previousWithStep()

	return id == next || id == previous
}

func (ring *Ring) next() int {
    return ring.peerID() + 1
}

func (ring *Ring) nextWithStep() int {
    return (ring.peerID() + ring.steps) % ring.peer.GetNumberParticipants()
}

func (ring *Ring) previous() int {
    return ring.previousNSteps(1)
}

func (ring *Ring) previousWithStep() int {
    return ring.previousNSteps(ring.steps)
}

func (ring *Ring) previousNSteps(steps int) int {
    peerID := ring.peerID()

    if (peerID - steps) <= 0 {
        return (ring.peer.GetNumberParticipants() + peerID) - steps
    }

    return peerID - steps
}

func (ring *Ring) peerID() int {
    return ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()
}
