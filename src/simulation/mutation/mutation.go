package mutation

import (
	con "simulation/consensus"
	ex "simulation/exception"
	stb "simulation/stubborn"
	tg "github.com/galeone/tfgo"
)

// kind of enums
const (
	EARLY       = iota
	RING        = iota
	GOSSIP      = iota
	CENTRALIZED = iota
	OLDRING     = iota
	ADAPTED	    = iota
)

// Mutation is an interface
type Mutation interface {
	Delta0(int, *stb.Package) bool
	Delta(int) bool
}

// NewMutation creates a new Mutation.
func NewMutation(peer *con.Peer, mutationType int, models ...*tg.Model) Mutation {
	switch mutationType {
	case EARLY:
		return NewEarly(peer)
	case RING:
		return NewRing(peer)
	case GOSSIP:
		return NewGossip(peer)
	case CENTRALIZED:
		return NewCentralized(peer)
	case OLDRING:
		return NewOldRing(peer)
	case ADAPTED:
		return NewAdapted(peer, models[0])
	}

	return NewEarly(peer)
}

// Find returns the constant number for a mutation.
func Find(mutation string) (int, error) {
	switch mutation {
	case "early":
		return EARLY, nil
	case "ring":
		return RING, nil
	case "gossip":
		return GOSSIP, nil
	case "centralized":
		return CENTRALIZED, nil
	case "old_ring":
		return OLDRING, nil
	case "adapted":
		return ADAPTED, nil
	}

	return -1, ex.NewError("Wrong mutation!")
}
