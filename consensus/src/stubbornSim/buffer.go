package stubborn

import (
	"log"
	"encoding/json"
	"sync"
	ex "exception"
)

// Buffer is a interface to absract the buffer implementation
type Buffer interface {
	insertElem(int, *Package)
	getElem(int)			  *Package
}

// BufferStruct is a struct where we keep the messages
type BufferStruct struct {
	OutBuffer	chan *Package
	InBuffer	chan *Package
	MutexIn 	*sync.Mutex
	MutexOut    *sync.Mutex
}

// Package is a struct to represent a package received via UDP.
type Package struct {
	ID	 	int
	Data	[]byte
	Arrived bool
	IsACK	bool
}

func newBuffer(size int) (buffer *BufferStruct) {
	buffer       = new(BufferStruct)
	buffer.Size  = size
	buffer.Data	 = make(map[int] *Package, size)
	buffer.Mutex = new(sync.Mutex)

	return
}

func newPackage(id int, data []byte, isAck bool) (pack *Package) {
	pack 	     = new(Package)
	pack.ID      = id 
	pack.Data    = data
	pack.IsACK	 = isAck
	pack.Arrived = false

	return 
}

func (b *BufferStruct) insertElem(id int, p *Package) {
	b.Mutex.Lock()
	b.Data[id] = p 
	b.Mutex.Unlock()
}

func (b BufferStruct) getElem(id int) *Package {
	elem, prs := b.Data[id]
	
	if prs {
		return elem
	}
	
	return nil
}

// GetID returns the peer ID that sends the package
func (p Package) GetID() int {
	return p.ID
}

// GetData returns the data in the Package
func (p Package) GetData() []byte {
	return p.Data
}

// PrintPacket prints all information in the Package
func (p Package) PrintPacket() {
	log.Println("-----------")	
	log.Print("ID: ")
	log.Println(p.ID)
	log.Print("Arrived: ")
	log.Println(p.Arrived)
	log.Print("isACK: ")
	log.Println(p.IsACK)
	log.Println("Data: ")
	log.Println(p.Data)	
	log.Println("-----------")
}

// Auxiliary Funtions 

func bytesToPackage(data []byte) (pack *Package) {
	err := json.Unmarshal(data, &pack)
	ex.CheckError(err)

	return
}