CURDIR :=$(shell pwd)
GOPATH :=$(CURDIR)
PATH   :=$(CURDIR)/bin:$(PATH)

ifeq (false,${DEBUG})
	export GOTRACEBACK=none
endif

compile:
	go install ./src/consensus/exception/
	go install ./src/consensus/message/
	go install ./src/consensus/stubborn/
	go install ./src/consensus/peer/

simulator:
	go install ./src/simulation/exception/
	go install ./src/simulation/consensus/
	go install ./src/simulation/stubborn/
	go install ./src/simulation/simulator/
	go install ./src/simulation/mutation/
	go install ./src/simulation/failuredetection/

install:
	go get "github.com/orcaman/concurrent-map"
	go get "github.com/joa/failuredetector"
	go get "github.com/yangwenmai/ratelimit/leakybucket"
	go get "github.com/yangwenmai/ratelimit/simpleratelimit"
	go get "github.com/kr/pretty"
	go get "github.com/tensorflow/tensorflow/tensorflow/go"
	go get "github.com/galeone/tfgo"
	go get "github.com/go-delve/delve/cmd/dlv"
	go get "go.uber.org/ratelimit"
	go get gonum.org/v1/plot/...

run_mutation:
	./bin/simulator ${MUTATION} ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${LATENCY} ${BANDWIDTH} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}

test:
	./bin/simulator early ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${LATENCY} ${BANDWIDTH} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulator ring ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${LATENCY} ${BANDWIDTH} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulator centralized ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${LATENCY} ${BANDWIDTH} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}
	./bin/simulator gossip ${NODES} ${DEFAULT_DELTA} ${MAX_TRIES} ${PERCENTAGE_MISS} ${LATENCY} ${BANDWIDTH} ${PERCENTAGE_FAULTS} ${PROBABILITY_TO_FAIL} ${WITH_ALL_METRICS}

format:
	go fmt ./src/simulation/exception/
	go fmt ./src/simulation/consensus/
	go fmt ./src/simulation/stubborn/
	go fmt ./src/simulation/simulator/
	go fmt ./src/simulation/mutation/
	go fmt ./src/simulation/failuredetection/

debug:
	dlv debug ./src/simulation/simulator/

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
