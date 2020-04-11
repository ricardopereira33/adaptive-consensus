package machinelearning

import (
	"strconv"

	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
)

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
				worker.model.Op("output_layer/add", 0),
			}, map[tf.Output]*tf.Tensor{
				worker.model.Op("input_layer_input", 0): tensor,
			},
		)
		request.responseChannel <- result[0].Value().([][][]float32)

		done <- worker
	}
}

func selectModel(numberParticipants int) string {
	return "mut_model_" + strconv.Itoa(numberParticipants)
}