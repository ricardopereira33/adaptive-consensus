CURDIR :=$(shell pwd)
GOPATH :=$(CURDIR)/consensus
PATH   :=$(CURDIR)/consensus/bin:$(PATH)

compile:
	go install ./consensus/src/stubborn/
	go install ./consensus/src/peer/

run:
	./consensus/bin/peer

