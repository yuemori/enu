.PHONY: test
test:
	make -B codegen
	go test -v

.PHONY: codegen
codegen: enumerable.generated.go comparer_enumerable.generated.go map_enumerable.generated.go ordered_enumerable.generated.go numeric_enumerable.generated.go

enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T any" "" "T"

comparer_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T comparable" "Comparer" "T"

ordered_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T constraints.Ordered" "Ordered" "T" '"golang.org/x/exp/constraints"'

numeric_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T constraints.Integer | constraints.Float" "Numeric" "T" '"golang.org/x/exp/constraints"'

map_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "K, V" "K comparable, V any" "Map" "KeyValuePair[K, V]"
