package stubborn

import (
	"log"
	"fmt"
	"encoding/json"
	"sync"
)

// Buffer is a interface to absract the buffer implementation
type Buffer interface {
	insertElem(*Package)
	getElem(int)		*Package
	getSize()			int
	getData()			map[int] *Package
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

func bytesToPackage(data []byte) (pack *Package) {
	err := json.Unmarshal(data, &pack)
	checkError(err, false)

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

func (b BufferStruct) getSize() int {
	return b.size
}

func (b BufferStruct) getData() map[int] *Package {
	return b.data
}

func (b BufferStruct) printBuffer() {
	log.Println("-----------")	
	log.Println("Size: " + fmt.Sprint(b.size))
	log.Println(b.data)	
	log.Println("-----------")
}