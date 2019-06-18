package mutation

import (
	"time"
	"strconv"
	// "math"

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
	consensusInfo := adapted.peer.GetConsensusInfo()

	if consensusInfo.Voters != nil && adapted.lastRequest != nil {
		value := adapted.lastRequest[id - 1]
		// value := math.Round(float64(adapted.lastRequest[id - 1]))

		if value > 0 {
			time.Sleep(time.Duration(value) * time.Millisecond)

			return true
		}
	}

	return false
}

// CacheQuerie caches the last query to the ML model
func (adapted *Adapted) CacheQuerie() {
	consensusInfo := adapted.peer.GetConsensusInfo()
	inputData := adapted.getConsensusStatus(consensusInfo)
	adapted.lastRequest = adapted.model.CreateRequest(inputData)[0][0]
}

func (adapted *Adapted) getConsensusStatus(consensus *con.Info) ([][][]float32) {
	finalList := make([][][]float32, 0)
	intermateList:= make([][]float32, 0)
	numberParticipants := adapted.peer.GetNumberParticipants()

	listOfVoters := getVoters(consensus.Voters, numberParticipants)

	if listOfVoters != nil {
		coordIDValues := getGenericValues(consensus.CoordID, numberParticipants)
		peerIDValues := getGenericValues(consensus.PeerID, numberParticipants)
		EstimatePeerIDValues := getGenericValues(consensus.Estimate.PeerID, numberParticipants)
		phaseValues := getPhaseValues(consensus.Phase)

		consensusStatus := []float32 { float32(consensus.Round) }

		consensusStatus = append(consensusStatus, listOfVoters...)
		consensusStatus = append(consensusStatus, normalizeDecision(consensus.Estimate.Value), normalizeDecision(consensus.Decision))
		consensusStatus = append(consensusStatus, peerIDValues...)
		consensusStatus = append(consensusStatus, coordIDValues...)
		consensusStatus = append(consensusStatus, EstimatePeerIDValues...)
		consensusStatus = append(consensusStatus, phaseValues...)

		intermateList = append(intermateList, consensusStatus)
		finalList = append(finalList, intermateList)
	}

	return finalList
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
				newList = append(newList, 1.0)
			} else {
				newList = append(newList, 0.0)
			}
		}

		return newList
	}

	return nil
}

func getGenericValues(peerID int, NumberParticipants int) []float32 {
	list := make([]float32, 0)

	for id := 1; id <= NumberParticipants; id++ {
		if id == peerID {
			list = append(list, 1.0)
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