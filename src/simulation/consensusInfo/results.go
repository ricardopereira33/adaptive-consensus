package consensusinfo

import "os"

// Results is the struct that contains the consensus results
type Results struct {
    PeerID int
    Info   map[string] string
}

// NewResults creates a new estimate
func NewResults(id int, info map[string] string) (results *Results) {
    results        = new(Results)
    results.PeerID = id
    results.Info   = info

    return
}

// Write writes all results into the file
func (r *Results) Write(file *os.File) {
    for _, value := range r.Info {
        file.WriteString(value)
    }

    file.WriteString("\n")
}