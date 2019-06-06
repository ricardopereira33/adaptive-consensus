package mutation

import (
	"time"
	"strconv"

	con "simulation/consensus"
	stb "simulation/stubborn"
	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
	cmap "github.com/orcaman/concurrent-map"
)

// Adapted is a mutation type
type Adapted struct {
	peer  *con.Peer
	model *tg.Model
}

// NewAdapted creates a new adapted mutation
func NewAdapted(peer *con.Peer) (adapted *Adapted) {
	adapted = new(Adapted)
	adapted.peer = peer
	adapted.model = newModel()

	return
}

// Delta0 is the delta0 implementation
func (adapted *Adapted) Delta0(id int, pack *stb.Package) bool {
	return adapted.Delta(id)
}

// Delta is the delta implementation
func (adapted *Adapted) Delta(id int) bool {
	inputData, voters := adapted.getConsensusStatus()

	if voters != nil {
		tensor, err := tf.NewTensor(inputData)
		ex.CheckError(err)

		result := adapted.model.Exec(
			[]tf.Output{
				adapted.model.Op("output_layer_1/add", 0),
			}, map[tf.Output]*tf.Tensor{
				adapted.model.Op("input_layer_input_1", 0): tensor,
			},
		)

		ex.CheckError(err)

		matrixValue := result[0].Value()
		delayValues := matrixValue.([][][]float32)[0][0]
		value := delayValues[id - 1]

		if value > 0 {
			time.Sleep(time.Duration(value) * time.Millisecond)

			return true
		}
	}

	return false
}

func (adapted *Adapted) getConsensusStatus() ([][][]float32, []float32) {
	consensus := adapted.peer.GetConsensusInfo()
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

	return finalList, listOfVoters
}

func newModel() *tg.Model {
	model := tg.LoadModel("src/simulation/models/mut_model", []string{"mut_tag"}, nil)

	return model
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