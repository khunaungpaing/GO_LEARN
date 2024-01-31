build:
	@go build -o bin/golearn
run: build
	@./bin/golearn
test:
	@go test ./...