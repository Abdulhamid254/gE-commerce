build:
	 @go build -o bin/gggcommerce

run: build
	@./bin/gggcommerce

test:
	@go test -v ./...