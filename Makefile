## build: builds the api into a executable file in the build directory
build:
	@go build -o build/api cmd/main.go

## test: runs defined test suite
test:
	@go test -v ./...

## run: builds the application then runs it
run: build
	@./build/api

## loadtest: uses the vegeta cli utility to loadtest the api
loadtest:
	@echo "GET http://localhost:3000/health" | vegeta attack -duration=60s -rate=10000/s | vegeta report

## dos: uses the vegeta cli utility to perform DoS on the api to test its limits
dos:
	@echo "GET http://localhost:3000/health" | vegeta attack -duration=60s -rate=50000/s | vegeta report

## help: displays help
help:
	@echo " Choose a command:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'