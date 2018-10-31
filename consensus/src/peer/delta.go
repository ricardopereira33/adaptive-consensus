package main

import (
	"stubborn"
	"mutation"
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

func newMutation(mutationType int) Mutation {
	switch mutationType {
	case EARLY:
		return mutation.NewEarly()
	case RING:
		return mutation.NewRing()
	case GOSSIP:
		return mutation.NewGossip()
	case CENTRALIZED:
		return mutation.NewCentralized()
	}

	return mutation.NewEarly()
}


