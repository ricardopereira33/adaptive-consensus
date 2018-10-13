package stubborn

import (
	"log"
	"fmt"
)

const debug = true

// Buffer is a interface to absract the buffer implementation
// FIFO 
type Buffer interface {
	insertElem([]byte) 	bool
	getElem()			[]byte
	getSize()			int
	getElements()		int
	getData()			[][]byte
	clear()
}

// BufferStruct is a struct where we keep the messages
type BufferStruct struct {
	elements int
	size 	 int
	data 	 [][]byte
}

func newBuffer(size int) (buffer *BufferStruct){
	buffer = new(BufferStruct)

	buffer.elements = 0
	buffer.size 	= size
	buffer.data		= make([][]byte, size)
	return
}

func (b BufferStruct) insertElem(elem []byte) bool {
	if b.elements == b.size{
		b.increase()
	}
	
	b.data = append(b.data, elem)
	b.elements++
	return true
}

func (b BufferStruct) getElem() []byte {
	if b.elements < (b.size/2){
		b.decrease()
	}

	elem   := b.data[0]
	b.data = b.data[1:]
	b.elements--
	return elem
}

func (b BufferStruct) increase() {
	tmpBuffer := make([][]byte, b.size*2)
	copy(tmpBuffer, b.data)
	b.data = tmpBuffer
	b.size *= 2
}

func (b BufferStruct) decrease() {
	tmpBuffer := make([][]byte, b.size/2)
	copy(tmpBuffer, b.data)
	b.data = tmpBuffer
	b.size /= 2
}

func (b BufferStruct) clear() {
	b.data = nil
	b.data = make([][]byte, 1)
}

func (b BufferStruct) getSize() int {
	return b.size
}

func (b BufferStruct) getElements() int {
	return b.elements
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