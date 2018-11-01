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

func NewMutation(mutationType int) Mutation {
	switch mutationType {
	case EARLY:
		return NewEarly()
	case RING:
		return NewRing()
	case GOSSIP:
		return NewGossip()
	case CENTRALIZED:
		return NewCentralized()
	}

	return NewEarly()
}


