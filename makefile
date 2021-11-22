BIN = 'gownload'
.PHONY: test fmt build test version

clean:	
	go clean -i ./...
	rm -f $(BIN)
fmt:
	go fmt ./...
build:
	make fmt
	go build -o $(BIN) main.go 
test:
	go test -v cmd/cmd_test.go -short
	go test -v util/util_test.go -short