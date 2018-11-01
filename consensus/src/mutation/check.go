package mutation

import (
	"stubborn"
	msg "message"
)

func fresh(old_pack *stubborn.Package, new_pack *stubborn.Package) bool {
	if old_pack != nil {
		old_msg := msg.PackageToMessage(old_pack)
		new_msg := msg.PackageToMessage(new_pack)
		
		sameRound := old_msg.Round == new_msg.Round
		samePhase := old_msg.Phase == new_msg.Round
		
		if sameRound && samePhase {
			return false
		}
	}
	
	return true 
}

func majority(pack *stubborn.Package, nPeers int) bool {
	message := msg.PackageToMessage(pack)

	if len(message.Voters) > nPeers/2 {
		return true
	}

	return false
}