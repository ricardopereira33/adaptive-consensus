CURDIR :=$(shell pwd)
GOPATH :=$(CURDIR)
PATH   :=$(CURDIR)/bin:$(PATH)

compile:
	go install ./src/consensus/exception/
	go install ./src/consensus/message/
	go install ./src/consensus/stubborn/
	go install ./src/consensus/peer/

simulation:
	go install ./src/simulation/exception/
	go install ./src/simulation/consensusInfo/
	go install ./src/simulation/message/
	go install ./src/simulation/stubborn/
	go install ./src/simulation/simulation/

install:
	go get "github.com/orcaman/concurrent-map"

run:
	./bin/peer

clean:
	rm -rf ./bin
	rm -rf ./pkg