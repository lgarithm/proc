default: .PHONEY

.PHONEY examples:
	go mod tidy
	GOBIN=$(CURDIR)/bin go install -v ./examples/...

run:
	PATH=$(CURDIR)/bin:$(PATH) \
		run-proc
