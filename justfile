default:
    @just --list

build:
    go build -o bin/app .

run:
    go run .

test:
    go test -v -coverpkg=. -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

integration-test:                
    go test -v -tags=integration -coverpkg=. -coverprofile=coverage.integration.out ./tests
    go tool cover -html=coverage.integration.out -o coverage.integration.html

benchmark:                                        
    go test -bench=. ./... | tee benchmark.out

fmt:
    go fmt ./...

vet:
    go vet ./...

check: fmt vet test

clean:
    rm -rf bin/
    rm -f coverage.out coverage.html benchmark.out
    go clean
