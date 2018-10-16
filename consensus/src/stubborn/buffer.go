package stubborn

import (
	"log"
	"encoding/json"
	"sync"
)

// Buffer is a interface to absract the buffer implementation
type Buffer interface {
	insertElem(*Package)
	getElem(int)		*Package
}

// BufferStruct is a struct where we keep the messages
type BufferStruct struct {
	size	int
	data 	map[int] *Package
	mutex 	*sync.Mutex
}

// Package is a struct to represent a package received via UDP.
type Package struct {
	id	 	int
	data	[]byte
	arrived bool
	isACK	bool
}

func newBuffer(size int) (buffer *BufferStruct) {
	buffer       = new(BufferStruct)
	buffer.size  = size
	buffer.data	 = make(map[int] *Package, size)
	buffer.mutex = new(sync.Mutex)

	return
}

func newPackage(id int, data []byte, isAck bool) (pack *Package) {
	pack 	     = new(Package)
	pack.id      = id
	pack.data    = data
	pack.isACK	 = isAck
	pack.arrived = false

	return 
}

func (b BufferStruct) insertElem(p *Package) {
	b.mutex.Lock()
	b.data[p.id-1] = p 
	b.mutex.Unlock()
}

func (b BufferStruct) getElem(id int) *Package {
	b.mutex.Lock()	
	elem, prs := b.data[id-1]
	b.mutex.Unlock()
	
	if prs {
		return elem
	}
	
	return nil
}

// GetID returns the peer ID that sends the package
func (p Package) GetID() int {
	return p.id
}

// GetData returns the data in the Package
func (p Package) GetData() []byte {
	return p.data
}

// PrintPacket prints all information in the Package
func (p Package) PrintPacket() {
	log.Println("-----------")	
	log.Print("ID: ")
	log.Println(p.id)
	log.Print("Arrived: ")
	log.Println(p.arrived)
	log.Print("isACK: ")
	log.Println(p.isACK)
	log.Println("Data: ")
	log.Println(p.data)	
	log.Println("-----------")
}

// Auxiliary Funtions 

func bytesToPackage(data []byte) (pack *Package) {
	err := json.Unmarshal(data, &pack)
	checkError(err, false)

	return
}