package machinelearning

import (
	"strconv"

	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
)

var modelSettings = map[string] map[string] string{
	"mut_model": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_old_40": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_20": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_merge_40": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_centralized_40": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_ring_40": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
}

// Worker is the struct that contains all attributes for a worker
type Worker struct {
	idx       int
	work      chan Request
	pending   int
	modelName string
	model     *tg.Model
}

func newWorker(work chan Request, numberParticipants int) (worker *Worker) {
	worker = new(Worker)
	worker.work = work
	worker.modelName = selectModel(numberParticipants)
	// worker.model = tg.LoadModel("src/simulation/models/" + worker.modelName + "_33", []string{"mut_tag"}, nil)
	worker.model = tg.LoadModel("src/simulation/models/" + worker.modelName + "_v1", []string{"mut_tag"}, nil)
	// worker.model = tg.LoadModel("src/simulation/models/" + worker.modelName + "_test_2", []string{"mut_tag"}, nil)
	worker.pending = 0

	return
}

func (worker *Worker) doWork(done chan *Worker) {
	for {
		request := <- worker.work
		tensor, err := tf.NewTensor(request.data)
		ex.CheckError(err)

		result := worker.model.Exec(
			[]tf.Output{
				worker.model.Op(modelSettings[worker.modelName]["output_layer"], 0),
			}, map[tf.Output]*tf.Tensor{
				worker.model.Op(modelSettings[worker.modelName]["input_layer"], 0): tensor,
			},
		)
		request.responseChannel <- result[0].Value().([][][]float32)

		done <- worker
	}
}

func selectModel(numberParticipants int) string {
	// return "mut_model_centralized_" + strconv.Itoa(numberParticipants)
	return "mut_model_merge_" + strconv.Itoa(numberParticipants)
}