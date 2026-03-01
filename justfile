set shell := ["bash", "-cu"]
set dotenv-load := true

default:
    @just --list

BINARY := "termuru"
OUTDIR := "bin"

# Run the application
run *args:
    go run ./... {{args}}

# Run go build to compile the code and output the binary to bin
build: mkdir
    go build -o {{OUTDIR}}/{{BINARY}} ./...

# Run gofmt to format the code in place
fmt:
    gofmt -w .

# Run go vet to check for potential issues in the code
vet:
    go vet ./...

# Run golangci-lint to perform linting checks on the code
lint:
    golangci-lint run ./...

# Run go test to execute all tests in the project
test:
    go test ./...

# Run all checks: fmt, vet, test and lint
check: fmt vet lint test

# Clean up the output directory by removing the files in the bin
clean:
    rm -rf {{OUTDIR}}

# Create the output directory if it doesn't exist
mkdir:
    mkdir -p {{OUTDIR}}
