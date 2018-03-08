.PHONY: fmt vet build clean test bench cover profiling release

COVERPROFILE=
DEBUG=
ARGS=

build: fmt vet test
	go build -o cup

fmt:
	go fmt ./...

vet:
	go vet -v .

clean:
	rm -f cup

test:
	go test -cover $(COVERPROFILE) -race $(DEBUG) $(ARGS) ./...

bench:
	go test -bench=. -benchmem $(ARGS) ./...

cover:
	$(eval COVERPROFILE += -coverprofile=coverage.out)
	go test -cover $(COVERPROFILE) -race $(ARGS) $(DEBUG)
	go tool cover -html=coverage.out
	rm -f coverage.out

profiling:
	go test -bench=. -cpuprofile cpu.out -memprofile mem.out $(ARGS)

release: clean fmt vet test
	GOOS=linux GOARCH=amd64 go build -o cup.linux_amd64
