package consensusinfo

// import "os"

// Results is the struct that contains the consensus results
type Results struct {
	PeerID   int
	Sent     []float64
	Received []float64
}

// NewResults creates a new estimate
func NewResults(sent, received []float64, id int) (results *Results) {
	results = new(Results)
	results.PeerID = id
	results.Sent = sent
	results.Received = received

	return
}

// Write writes all results into the file
// func (r *Results) Write(file *os.File) {
//     for _, value := range r.Info {
//         file.WriteString(value)
//     }

//     file.WriteString("\n")
// }
