.PHONY: test
test:
	make -B codegen
	go test -v

.PHONY: codegen
codegen: enumerator.generated.go enumerator_comparable.generated.go enumerator_map.generated.go enumerator_ordered.generated.go enumerator_numeric.generated.go

enumerator.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T any" "" "T"

enumerator_comparable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T comparable" "Comparable" "T"

enumerator_ordered.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T constraints.Ordered" "Ordered" "T" '"golang.org/x/exp/constraints"'

enumerator_numeric.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "T" "T constraints.Integer | constraints.Float" "Numeric" "T" '"golang.org/x/exp/constraints"'

enumerator_map.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerator.go.tpl $@ "K, V" "K comparable, V any" "Map" "KeyValuePair[K, V]"
