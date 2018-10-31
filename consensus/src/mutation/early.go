package mutation

import "stubborn"

type Early struct { }

func NewEarly() *Early {
	return new(Early)
}

func (e Early) Delta0(id int, message *stubborn.Package) bool {
	return true
}
	
func (e Early) Delta(id int) bool {
	return true
}