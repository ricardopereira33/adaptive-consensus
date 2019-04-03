package stubborn

import (
	"encoding/json"
	ex "simulation/exception"
	"sync"
)

// Buffer is a interface to absract the buffer implementation
type Buffer interface {
	InsertElement(int, *Package)
	GetElement(int) *Package
}

// BufferStruct is a struct where we keep the messages
type BufferStruct struct {
	Size  int
	Data  map[int]*Package
	Mutex *sync.Mutex
}

// Package is a struct to represent a package received via UDP.
type Package struct {
	ID           int
	Data         []byte
	Arrived      bool
	Ack          bool
	Suspicious   bool
	SuspiciousID int
}

func newBuffer(size int) (buffer *BufferStruct) {
	buffer = new(BufferStruct)
	buffer.Size = size
	buffer.Data = make(map[int]*Package, size)
	buffer.Mutex = new(sync.Mutex)

	return
}

func newPackage(id int, data []byte, ack bool) (pack *Package) {
	pack = new(Package)
	pack.ID = id
	pack.Data = data
	pack.Ack = ack
	pack.Arrived = false

	return
}

func newSuspect(suspiciousID int, suspicious bool) (pack *Package) {
	pack = new(Package)
	pack.SuspiciousID = suspiciousID
	pack.Suspicious = suspicious

	return
}

/*** Exported methods ***/

// InsertElement insert a package to the buffer
func (buffer *BufferStruct) InsertElement(id int, pack *Package) {
	buffer.Mutex.Lock()
	buffer.Data[id] = pack
	buffer.Mutex.Unlock()
}

// GetElement get a package for the process "id"
func (buffer BufferStruct) GetElement(id int) *Package {
	buffer.Mutex.Lock()
	pack, present := buffer.Data[id]
	buffer.Mutex.Unlock()

	if present {
		return pack
	}

	return nil
}

// GetID returns the peer ID that sends the package
func (pack Package) GetID() int {
	return pack.ID
}

// GetData returns the data in the Package
func (pack Package) GetData() []byte {
	return pack.Data
}

/*** Auxiliary Funtions ***/

func bytesToPackage(data []byte) (pack *Package) {
	err := json.Unmarshal(data, &pack)
	ex.CheckError(err)

	return
}
