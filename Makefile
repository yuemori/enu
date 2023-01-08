.PHONY: test
test:
	go test -v

.PHONY: codegen
codegen: enumerator.generated.go enumerator_comparable.generated.go enumerator_map.generated.go

enumerator.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T any" "" "T"

enumerator_comparable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T comparable" "C" "T"

enumerator_map.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "K, V" "K comparable, V any" "Map" "KeyValuePair[K, V]"
