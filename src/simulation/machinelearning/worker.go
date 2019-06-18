package machinelearning

import (
	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
)

const (
	model40 = "mut_model_40"
	model80 = "mut_model_80"
)

var modelSettings = map[string] map[string] string{
	"mut_model_40": map[string] string {
		"input_layer"  : "input_layer_input",
		"output_layer" : "output_layer/add",
	},
	"mut_model_80": map[string] string {
		"input_layer"  : "input_layer_input_1",
		"output_layer" : "output_layer_1/add",
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
	worker.model = tg.LoadModel("src/simulation/models/" + worker.modelName, []string{"mut_tag"}, nil)
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
	switch numberParticipants {
	case 40:
		return model40
	case 80:
		return model80
	}

	return model40
}