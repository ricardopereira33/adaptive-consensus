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
	go install ./src/simulation/mutation/

test:
	./bin/simulation early 50 3 3 1
	./bin/simulation ring 50 3 3 1
	./bin/simulation centralized 50 3 3 1
	./bin/simulation gossip 50 3 3 1

format:
	go fmt ./src/simulation/exception/
	go fmt ./src/simulation/consensusInfo/
	go fmt ./src/simulation/message/
	go fmt ./src/simulation/stubborn/
	go fmt ./src/simulation/simulation/
	go fmt ./src/simulation/mutation/

install:
	go get "github.com/orcaman/concurrent-map"
	go get gonum.org/v1/plot/...

run:
	./bin/peer

clean:
	rm -rf ./bin
	rm -rf ./pkg
	rm *.png