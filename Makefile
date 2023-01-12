.PHONY: test
test:
	make -B codegen
	go test -v

.PHONY: codegen
codegen: Makefile enumerable.generated.go comparer_enumerable.generated.go map_enumerable.generated.go ordered_enumerable.generated.go numeric_enumerable.generated.go

enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T any" "" "T" ""

comparer_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T comparable" "Comparer" "T" "" "Uniq,Contains,IndexOf,ToMap"

ordered_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T constraints.Ordered" "Ordered" "T" '"golang.org/x/exp/constraints"' "ToMap,Min,Max,Sort,Uniq,Contains,IndexOf"

numeric_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "T" "T constraints.Integer | constraints.Float" "Numeric" "T" '"golang.org/x/exp/constraints"' "Uniq,Min,Max,Sum,Sort,Contains,IndexOf,ToMap"

map_enumerable.generated.go: $(wildcard templates/*)
	go run templates/main.go templates/enumerable.go.tpl $@ "K, V" "K comparable, V any" "Map" "KeyValuePair[K, V]" ""
