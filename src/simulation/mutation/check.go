package mutation

import (
	msg "simulation/message"
	"simulation/stubborn"
)

func fresh(oldPack *stubborn.Package, newPack *stubborn.Package) bool {
	if oldPack != nil {
		oldMsg := msg.PackageToMessage(oldPack)
		newMsg := msg.PackageToMessage(newPack)

		sameRound := oldMsg.Round == newMsg.Round
		samePhase := oldMsg.Phase == newMsg.Round

		if sameRound && samePhase {
			return false
		}
	}

	return true
}

func majority(pack *stubborn.Package, numberParticipants int) bool {
	message := msg.PackageToMessage(pack)

	if len(message.Voters) > numberParticipants/2 {
		return true
	}

	return false
}
