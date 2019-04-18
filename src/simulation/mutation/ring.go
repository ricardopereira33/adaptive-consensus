package mutation

import (
	con "simulation/consensus"
    stb "simulation/stubborn"
)

// Ring is a mutation of a ring
type Ring struct {
    peer                  *con.Peer
    stepsForward          int
    stepsToBack           int
    suspectFailureForward bool
    suspectFailureToBack  bool
}

// NewRing creates a new ring
func NewRing(peer *con.Peer) (ring *Ring) {
	ring = new(Ring)
    ring.peer = peer
    ring.stepsForward = 1
    ring.stepsToBack = 1
    ring.suspectFailureForward = false
    ring.suspectFailureToBack = false

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
    next := ring.next()
    previous := ring.previous()

    lastPackage := ring.peer.GetChannel().LastPackageBuffered(id)
    packageArrived := lastPackage != nil && lastPackage.Arrived

    if !packageArrived {
        if id == next && !ring.suspectFailureForward {
            ring.stepsForward++
            ring.suspectFailureForward = true
        } else if id == previous && !ring.suspectFailureToBack {
            ring.stepsToBack++
            ring.suspectFailureToBack = true
        }
    }

    if ring.isLastNode(id) {
        ring.suspectFailureForward = false
        ring.suspectFailureToBack = false
    }

	return id == next || id == previous
}

func (ring *Ring) next() int {
    nextID := ring.peerID() + ring.stepsForward
    numberParticipants := ring.peer.GetNumberParticipants()

    if nextID > numberParticipants {
        return nextID % numberParticipants
    }

    return nextID
}

func (ring *Ring) previous() int {
    peerID := ring.peerID()
    steps := ring.stepsToBack

    if (peerID - steps) <= 0 {
        return (ring.peer.GetNumberParticipants() + peerID) - steps
    }

    return peerID - steps
}

func (ring *Ring) peerID() int {
    return ring.peer.GetPeerID() % ring.peer.GetNumberParticipants()
}

func (ring *Ring) isLastNode(dest int) bool {
    numberParticipants := ring.peer.GetNumberParticipants()

    if ring.peerID() == numberParticipants {
        return dest == numberParticipants - 1
    }

    return dest == numberParticipants
}