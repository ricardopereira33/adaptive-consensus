package mutation

import (
	"time"
	"strconv"

	con "simulation/consensus"
	stb "simulation/stubborn"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
	cmap "github.com/orcaman/concurrent-map"
)

// Adapted is a mutation type
type Adapted struct {
	peer  *con.Peer
	model *tf.SavedModel
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
	inputData := adapted.getConsensusStatus()
	tensor, err := tf.NewTensor(inputData)
	ex.CheckError(err)

	result, err := adapted.model.Session.Run(
		map[tf.Output]*tf.Tensor {
			adapted.model.Graph.Operation("input_layer_input_3").Output(0): tensor,
		},
		[]tf.Output{
			adapted.model.Graph.Operation("output_layer_3/BiasAdd").Output(0),
		},
		nil,
	)

	ex.CheckError(err)

	matrixValue := result[0].Value()
	value := matrixValue.([][][]float32)[0][0][id - 1]

	if value > 0 {
		time.Sleep(time.Duration(value) * time.Millisecond)

		return true
	}

	return false
}

func (adapted *Adapted) getConsensusStatus() [][][]float32 {
	consensus := adapted.peer.GetConsensusInfo()
	finalList := make([][][]float32, 0)
	intermateList:= make([][]float32, 0)

	consensusStatus := []float32 {
		float32(consensus.CoordID),
		float32(consensus.Round),
		float32(consensus.Phase),
		float32(consensus.Estimate.PeerID),
		normalizeDecision(consensus.Estimate.Value),
		normalizeDecision(consensus.Decision)}

	listOfVoters := getVoters(consensus.Voters, adapted.peer.GetNumberParticipants())

	consensusStatus = append(consensusStatus, listOfVoters...)
	intermateList = append(intermateList, consensusStatus)
	finalList = append(finalList, intermateList)

	return finalList
}

func newModel() *tf.SavedModel {
	model, err := tf.LoadSavedModel("src/simulation/models/mut_model", []string{"mut_tag"}, nil)
	ex.CheckError(err)

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

	for id := 0; id < NumberParticipants; id++ {
		present := voters.Has(strconv.Itoa(id))

		if present {
			newList = append(newList, 1.0)
		} else {
			newList = append(newList, 0.0)
		}
	}

	return newList
}
