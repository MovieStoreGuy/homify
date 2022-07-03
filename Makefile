build:
	CGO_ENABLED=0 go build \
		--ldflags="-s -w" \
		--trimpath \
		-o homify \
		./cmd/homify

lint:
	golangci-lint run --allow-parallel-runners ./...

test:
	go test --cover --race --count=1 ./...

clean:
	$(RM) homify

.PHONY: build test lint