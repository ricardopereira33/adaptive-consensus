package mutation

import "consensus/stubborn"

type Ring struct { 
	channel stubborn.StubChannel
}

func NewRing(channel stubborn.StubChannel) (r *Ring) {
	r 		  = new(Ring)
	r.channel = channel

	return r
}

func (r Ring) Delta0(id int, pack *stubborn.Package) bool {
	isFresh    := fresh(r.channel.GetPackage(id), pack)
	isMajority := majority(pack, r.channel.GetNParticipants())
	isIDEqual  := id == ((r.channel.GetPeerID() % r.channel.GetNParticipants()) + 1)
	
	return isIDEqual && (isFresh || isMajority)
}
	
func (r Ring) Delta(id int) bool {
	return id == ((r.channel.GetPeerID() % r.channel.GetNParticipants()) + 1)
}