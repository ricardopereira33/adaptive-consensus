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

install:
	go get "github.com/orcaman/concurrent-map"
	go get "github.com/joa/failuredetector"
	go get gonum.org/v1/plot/...
	go get "go.uber.org/ratelimit"

run_mutation:
	./bin/simulation ${MUTATION} ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${WITH_FAULTS} ${LATENCY} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}

test:
	./bin/simulation early ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${WITH_FAULTS} ${LATENCY} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulation ring ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${WITH_FAULTS} ${LATENCY} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulation centralized ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${WITH_FAULTS} ${LATENCY} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulation gossip ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${WITH_FAULTS} ${LATENCY} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}

format:
	go fmt ./src/simulation/exception/
	go fmt ./src/simulation/consensus/
	go fmt ./src/simulation/stubborn/
	go fmt ./src/simulation/simulation/
	go fmt ./src/simulation/mutation/
	go fmt ./src/simulation/failuredetection/

run:
	./bin/peer

clean_results:
	rm *.csv
	rm *.png
	rm -rf ./results

clean:
	rm -rf ./bin
	rm -rf ./pkg
	rm *.csv
	rm *.png
