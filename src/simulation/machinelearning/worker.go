package machinelearning

import (
	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	ex "simulation/exception"
)

// Worker is the struct that contains all attributes for a worker
type Worker struct {
	idx     int
	work    chan Request
	pending int
	model   *tg.Model
}

func newWorker(work chan Request) (worker *Worker) {
	worker = new(Worker)
	worker.work = work
	worker.model = tg.LoadModel("src/simulation/models/mut_model", []string{"mut_tag"}, nil)
	worker.pending = 0

	return
}

func (worker *Worker) doWork(done chan *Worker) {
	for {
		request := <- worker.work
		// log.Println("WORKING!!")
		tensor, err := tf.NewTensor(request.data)
		ex.CheckError(err)

		result := worker.model.Exec(
			[]tf.Output{
				worker.model.Op("output_layer_1/add", 0),
			}, map[tf.Output]*tf.Tensor{
				worker.model.Op("input_layer_input_1", 0): tensor,
			},
		)
		request.responseChannel <- result[0].Value().([][][]float32)
		// println("WORK SENT!")
		done <- worker
	}
}
