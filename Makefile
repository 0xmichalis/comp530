build:
	@go build ./...
.PHONY: build

profile:
	@go test -bench=. -count=1 ./...
.PHONY: profile
