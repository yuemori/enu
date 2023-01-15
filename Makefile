.PHONY: test
test:
	make -B codegen
	go test -v
	golangci-lint run

.PHONY: codegen
codegen:
	go run templates/main.go
