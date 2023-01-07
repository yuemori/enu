.PHONY: codegen
codegen: enumerator.generated.go enumerator_c.generated.go

enumerator.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl enumerator.generated.go "T" "T any" ""

enumerator_c.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl enumerator_c.generated.go "T" "T comparable" "C"
