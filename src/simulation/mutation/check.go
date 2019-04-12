package mutation

import (
	cons "simulation/consensus"
	stb "simulation/stubborn"
)

func fresh(oldPackage *stb.Package, newPackage *stb.Package) bool {
	if oldPackage != nil {
		oldMessage := cons.PackageToMessage(oldPackage)
		newMessage := cons.PackageToMessage(newPackage)

		sameRound := oldMessage.Round == newMessage.Round
		samePhase := oldMessage.Phase == newMessage.Round

		if sameRound && samePhase {
			return false
		}
	}

	return true
}

func majority(pack *stb.Package, numberParticipants int) bool {
	message := cons.PackageToMessage(pack)

	if len(message.Voters) > numberParticipants/2 {
		return true
	}

	return false
}
