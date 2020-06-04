default: .PHONEY

.PHONEY examples:
	go mod tidy
	GOBIN=$(CURDIR)/bin go install -v ./examples/...

test:
	go test -v ./...

run: .PHONEY
	PATH=$(CURDIR)/bin:$(PATH) \
		run-proc
