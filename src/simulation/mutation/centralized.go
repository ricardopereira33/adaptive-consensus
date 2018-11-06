package mutation

import "simulation/stubborn"

type Centralized struct {
	channel stubborn.StubChannel
}

func NewCentralized(channel stubborn.StubChannel) (c *Centralized) {
	c         = new(Centralized)
	c.channel = channel 

	return c
}

func (c Centralized) Delta0(id int, pack *stubborn.Package) bool {
	coordID := c.channel.GetCoordID()
	peerID  := c.channel.GetPeerID()

	isCoord    := id == coordID || peerID == coordID
	isFresh    := fresh(c.channel.GetPackage(id), pack)
	isMajority := majority(pack, c.channel.GetNParticipants())
	
	return isCoord && (isFresh || isMajority)
}
	
func (c Centralized) Delta(id int) bool {
	return true	
}