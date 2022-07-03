build:
	CGO_ENABLED=0 go build \
		--ldflags="-s -w" \
		--trimpath \
		-o homify \
		./cmd/homify


test:

clean:
	$(RM) homify

.PHONY: build test