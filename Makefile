default: .PHONEY

.PHONEY examples:
	go mod tidy
	GOBIN=$(CURDIR)/bin go install -v ./examples/...

test:
	go test -v ./...

run:
	PATH=$(CURDIR)/bin:$(PATH) \
		run-proc
