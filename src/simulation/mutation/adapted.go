package mutation

import (
	"time"
	"strconv"
	"math"
	"fmt"

	con "simulation/consensus"
	stb "simulation/stubborn"
	cmap "github.com/orcaman/concurrent-map"
	ml "simulation/machinelearning"
)

// Adapted is a mutation type
type Adapted struct {
	peer        *con.Peer
	model       *ml.Balancer
	lastRequest []float32
}

// NewAdapted creates a new adapted mutation
func NewAdapted(peer *con.Peer, model *ml.Balancer) (adapted *Adapted) {
	adapted = new(Adapted)
	adapted.peer = peer
	adapted.model = model
	adapted.lastRequest = nil

	return
}

// Delta0 is the delta0 implementation
func (adapted *Adapted) Delta0(id int, pack *stb.Package) bool {
	return adapted.Delta(id)
}

// Delta is the delta implementation
func (adapted *Adapted) Delta(id int) bool {
	if adapted.lastRequest != nil {
		value := math.Round(float64(adapted.lastRequest[id - 1]))

		if value >= 0 {
			time.Sleep(time.Duration(value) * time.Millisecond)

			return true
		}
	}

	return false
}

// CacheQuerie caches the last query to the ML model
func (adapted *Adapted) CacheQuerie(pack *stb.Package) {
	inputData := adapted.getConsensusStatus(pack)

	adapted.lastRequest = adapted.model.CreateRequest(inputData)[0][0]

	id := adapted.peer.GetPeerID()
	if id < 35 && id > 5 {
		// [0][0][:42]
		fmt.Println("PEER ", id, " | INPUT: ", inputData, "\nPREDITION: ", adapted.lastRequest[(id-2):(id+2)])
	} else if id <= 5 {
		fmt.Println("PEER ", id, " | INPUT: ", inputData, "\nPREDITION: ", adapted.lastRequest[38:], adapted.lastRequest[:(id+2)])
	} else {
		fmt.Println("PEER ", id, " | INPUT: ", inputData, "\nPREDITION: ", adapted.lastRequest[(id-2):], adapted.lastRequest[:2])
	}
}

func (adapted *Adapted) getConsensusStatus(pack *stb.Package) ([][][]float32) {
	consensus := adapted.peer.GetConsensusInfo()
	finalList := make([][][]float32, 0)
	intermediateList:= make([][]float32, 0)
	numberParticipants := adapted.peer.GetNumberParticipants()

	listOfVoters := getVoters(consensus.Voters, numberParticipants)

	if listOfVoters != nil {
		coordIDValues := getGenericValues(consensus.CoordID, numberParticipants, 1.0)
		peerIDValues := getGenericValues(consensus.PeerID, numberParticipants, 200.0)
		// EstimatePeerIDValues := getGenericValues(consensus.Estimate.PeerID, numberParticipants)
		// phaseValues := getPhaseValues(consensus.Phase)

		// consensusStatus := []float32 { float32(consensus.Round), isMajority(pack, numberParticipants) }
		consensusStatus := []float32 { isMajority(pack, numberParticipants), normalizeBandwidth(adapted.peer.GetChannel().GetBandwidth()) }

		consensusStatus = append(consensusStatus, listOfVoters...)
		// consensusStatus = append(consensusStatus, getIsFreshValues(adapted.peer.GetChannel(), pack, numberParticipants)...)
		// consensusStatus = append(consensusStatus, normalizeDecision(consensus.Estimate.Value), normalizeDecision(consensus.Decision))
		consensusStatus = append(consensusStatus, normalizeDecision(consensus.Decision))
		consensusStatus = append(consensusStatus, peerIDValues...)
		consensusStatus = append(consensusStatus, coordIDValues...)
		// consensusStatus = append(consensusStatus, EstimatePeerIDValues...)
		// consensusStatus = append(consensusStatus, phaseValues...)

		intermediateList = append(intermediateList, consensusStatus)
		finalList = append(finalList, intermediateList)
	}

	return finalList
}

func normalizeBandwidth(value int) float32 {
	// minValue := float32(0.0)
	// maxValue := float32(500.0)

	// return (float32(value) - minValue)/(maxValue - minValue)
	return float32(value)
}

func normalizeDecision(value string) float32 {
	if value != "" {
		return 1.0
	}

	return 0.0
}

func getVoters(voters cmap.ConcurrentMap, NumberParticipants int) []float32 {
	newList := make([]float32, 0)

	if voters != nil {
		for id := 1; id <= NumberParticipants; id++ {
			present := voters.Has(strconv.Itoa(id))

			if present {
				newList = append(newList, 20.0)
			} else {
				newList = append(newList, 0.0)
			}
		}

		return newList
	}

	return nil
}

func getGenericValues(peerID int, NumberParticipants int, value float32) []float32 {
	list := make([]float32, 0)

	for id := 1; id <= NumberParticipants; id++ {
		if id == peerID {
			list = append(list, value)
		} else {
			list = append(list, 0.0)
		}
	}

	return list
}

func getPhaseValues(phase int) []float32 {
	if phase == 1 {
		return []float32{ 1.0, 0.0 }
	}

	return []float32{ 0.0, 1.0 }
}

func getIsFreshValues(channel stb.SChannel, pack *stb.Package, NumberParticipants int) []float32 {
	list := make([]float32, 0)

	for id := 1; id <= NumberParticipants; id++ {
		if fresh(channel.GetPackage(id), pack) {
			list = append(list, 1.0)
		} else {
			list = append(list, 0.0)
		}
	}

	return list
}

func isMajority(pack *stb.Package, numberParticipants int) float32 {
	if majority(pack, numberParticipants) {
		return 10.0
	}

	return 0.0
}
