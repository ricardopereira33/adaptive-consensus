package mutation

import (
	"stubborn"
)

// kind of enums
const (
	EARLY 		= iota
	RING 		= iota
	GOSSIP 		= iota
	CENTRALIZED = iota
)

// Mutation is an interface
type Mutation interface {
	Delta0(int, *stubborn.Package) bool
	Delta(int) 					   bool
}

func NewMutation(channel stubborn.StubChannel, mutationType int) Mutation {
	switch mutationType {
	case EARLY:
		return NewEarly(channel)
	case RING:
		return NewRing(channel)
	case GOSSIP:
		return NewGossip(channel)
	case CENTRALIZED:
		return NewCentralized(channel)
	}

	return NewEarly(channel)
}


