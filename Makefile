CURDIR :=$(shell pwd)
GOPATH :=$(CURDIR)/consensus
PATH   :=$(CURDIR)/consensus/bin:$(PATH)

compile:
	go install ./consensus/src/errors/
	go install ./consensus/src/message/
	go install ./consensus/src/stubborn/
	go install ./consensus/src/peer/

run:
	./consensus/bin/peer

clean:
	rm -rf ./consensus/bin
	rm -rf ./consensus/pkg

