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
	go install ./src/simulation/consensus/
	go install ./src/simulation/stubborn/
	go install ./src/simulation/simulation/
	go install ./src/simulation/mutation/
	go install ./src/simulation/failuredetection/

test:
	./bin/simulation early 100 3 3 1
	./bin/simulation ring 100 3 3 1
	./bin/simulation centralized 100 3 3 1
	./bin/simulation gossip 100 3 3 1

format:
	go fmt ./src/simulation/exception/
	go fmt ./src/simulation/consensus/
	go fmt ./src/simulation/stubborn/
	go fmt ./src/simulation/simulation/
	go fmt ./src/simulation/mutation/
	go fmt ./src/simulation/failuredetection/

install:
	go get "github.com/orcaman/concurrent-map"
	go get "github.com/joa/failuredetector"
	go get gonum.org/v1/plot/...

run:
	./bin/peer

clean_results:
	rm *.csv
	rm *.png

clean:
	rm -rf ./bin
	rm -rf ./pkg
	rm *.csv
	rm *.png
