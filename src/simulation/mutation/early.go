package mutation

import "simulation/stubborn"

// Early is a mutation type
type Early struct {
	channel stubborn.StubChannel
}

// NewEarly creates a new early mutation
func NewEarly(channel stubborn.StubChannel) (early *Early) {
	early = new(Early)
	early.channel = channel

	return 
}

// Delta0 is the delta0 implementation
func (early Early) Delta0(id int, pack *stubborn.Package) bool {
	isFresh := fresh(early.channel.GetPackage(id), pack)
	isMajority := majority(pack, early.channel.GetNumberParticipants())

	return isFresh || isMajority
}

// Delta is the delta implementation
func (early Early) Delta(id int) bool {
	return true
}
