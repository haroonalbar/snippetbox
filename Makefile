build:
	@go build -o ./bin/snip

run: build
	@./bin/snip

test:
	@go test ./... -v
