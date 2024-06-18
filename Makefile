build:
	@go build -o ./bin/snip ./cmd/web

run: build
	@./bin/snip

test:
	@go test ./... -v
