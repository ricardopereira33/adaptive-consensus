package mutation

import (
	"math/rand"
	con "simulation/consensus"
	stb "simulation/stubborn"
)

// Gossip is a mutation type
type Gossip struct {
	peer            *con.Peer
	permut          []int
	turn            int
	fanout          int
	processesToSkip int
}

// NewGossip creates a new gossip mutation
func NewGossip(peer *con.Peer) (gossip *Gossip) {
	gossip = new(Gossip)
	gossip.peer = peer
	gossip.permut = perm(peer.GetNumberParticipants())
	gossip.processesToSkip = 1
	gossip.turn = 0
	gossip.fanout = 5

	return
}

// Delta0 is the delta0 implementation
func (gossip *Gossip) Delta0(id int, message *stb.Package) bool {
	needAck := gossip.peer.NeedAck(id)

	return gossip.Delta(id) || needAck
}

// Delta is the delta implementation
func (gossip *Gossip) Delta(id int) bool {
	numberParticipants := len(gossip.permut)
	gossip.turn++

	if gossip.turn == numberParticipants {
		gossip.processesToSkip += gossip.fanout
		gossip.turn = 0
	}

	processIndex := 0

	for processIndex < gossip.fanout {
		if gossip.permut[(processIndex + gossip.processesToSkip) % numberParticipants] == id {
			return true
		}
		processIndex++
	}

	return false
}

func perm(numberParticipants int) (list []int) {
	list = rand.Perm(numberParticipants)

	for index := range list {
		list[index]++
	}

	return
}
