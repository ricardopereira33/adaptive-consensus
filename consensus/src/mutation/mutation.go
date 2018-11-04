package mutation

import (
	"stubborn"
	ex "exception"
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

// NewMutation creates a new Mutation.
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
	}

	return -1, ex.NewError("Wrong mutation!")
}


