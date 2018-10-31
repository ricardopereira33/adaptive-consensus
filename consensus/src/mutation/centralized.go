package mutation

import "stubborn"

type Centralized struct { }

func NewCentralized() *Centralized {
	return new(Centralized)
}

func (c Centralized) Delta0(int, *stubborn.Package) bool {
	return true
}
	
func (c Centralized) Delta(int) bool {
	return true	
}