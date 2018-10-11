GOPATH :=$(HOME)/Documents/Git/Thesis/adaptive-consensus/consensus
PATH   :=$(HOME)/Documents/Git/Thesis/adaptive-consensus/consensus/bin:$(PATH)

compile:
	go install ./consensus/src/peer/

run:
	./consensus/bin/peer

