.PHONY: bench
bench:
	@go test -bench=. -benchmem ./...

.PHONY: ut
ut:
	@go test -tags=goexperiment.arenas -race ./...

.PHONY: fmt
fmt:
	@sh ./scripts/goimports.sh

.PHONY: lint
lint:
	@golangci-lint run -c .golangci.yml

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check: fmt tidy