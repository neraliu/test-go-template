GOPATH=$(shell pwd)
GOBIN=${GOPATH}/bin

build:
	@echo "####################################"
	@echo "# Building .... on ${PLATFORMS}"
	@echo "####################################"
	mkdir -p ${GOBIN}
	gofmt -w src/*.go
	GOPATH=${GOPATH} GOBIN=${GOBIN} go install src/template.go
