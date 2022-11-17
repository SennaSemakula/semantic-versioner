PHONY: gen

gen:
	@go generate

run: gen
	@go run .

test:
	go test -v ./...