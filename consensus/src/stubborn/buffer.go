package stubborn

import (
	"log"
	"fmt"
)

const debug = true

// Buffer is a interface to absract the buffer implementation
// FIFO 
type Buffer interface {
	insertElem(int, []byte) bool
	getElem(int)			[]byte
	getSize()				int
	getData()				[][]byte
}

// BufferStruct is a struct where we keep the messages
type BufferStruct struct {
	size 	 int
	data 	 [][]byte
}

func newBuffer(size int) (buffer *BufferStruct){
	buffer = new(BufferStruct)

	buffer.size 	= size
	buffer.data		= make([][]byte, size)
	return
}

func (b BufferStruct) insertElem(id int, elem []byte) bool {
	b.data[id-1] = elem
	return true
}

func (b BufferStruct) getElem(id int) []byte {
	elem := b.data[id-1]
	return elem
}

func (b BufferStruct) getSize() int {
	return b.size
}

func (b BufferStruct) getData() [][]byte {
	return b.data
}

func (b BufferStruct) printBuffer() {
	log.Println("-----------")	
	log.Println("Elems: " + fmt.Sprint(b.elements))
	log.Println("Size: " + fmt.Sprint(b.elements))
	log.Println(b.data)	
	log.Println("-----------")
}