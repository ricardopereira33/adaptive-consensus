package mutation

import "stubborn"

type Ring struct { }

func NewRing() *Ring {
	return new(Ring)
}

func (r Ring) Delta0(id int, message *stubborn.Package) bool {
	return true
}
	
func (r Ring) Delta(id int) bool {
	return true
}