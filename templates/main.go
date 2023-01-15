package main

import (
	"log"
	"os"
	"text/template"
)

var extrasMethods = map[string][]string{
	"ToMap":    {"", "Comparer", "Numeric", "Ordered"},
	"Uniq":     {"Comparer", "Numeric", "Ordered"},
	"Contains": {"Comparer", "Numeric", "Ordered"},
	"IndexOf":  {"Comparer", "Numeric", "Ordered"},
	"Max":      {"Numeric", "Ordered"},
	"Min":      {"Numeric", "Ordered"},
	"Sort":     {"Numeric", "Ordered"},
	"Sum":      {"Numeric"},
}

func main() {
	data := map[string]struct {
		Name               string
		Type               string
		TypeWithConstraint string
		Prefix             string
		ItemType           string
		ImportPkg          string
		Extras             []string
	}{
		"enumerable.generated.go": {
			Type:               "T",
			TypeWithConstraint: "T any",
			Prefix:             "",
			ItemType:           "T",
			ImportPkg:          "",
		},
		"comparer_enumerable.generated.go": {
			Type:               "T",
			TypeWithConstraint: "T comparable",
			Prefix:             "Comparer",
			ItemType:           "T",
			ImportPkg:          "",
		},
		"ordered_enumerable.generated.go": {
			Type:               "T",
			TypeWithConstraint: "T constraints.Ordered",
			Prefix:             "Ordered",
			ItemType:           "T",
			ImportPkg:          "golang.org/x/exp/constraints",
		},
		"numeric_enumerable.generated.go": {
			Type:               "T",
			TypeWithConstraint: "T constraints.Integer | constraints.Float",
			Prefix:             "Numeric",
			ItemType:           "T",
			ImportPkg:          "golang.org/x/exp/constraints",
		},
		"map_enumerable.generated.go": {
			Type:               "K, V",
			TypeWithConstraint: "K comparable, V any",
			Prefix:             "Map",
			ItemType:           "KeyValuePair[K, V]",
			ImportPkg:          "",
		},
	}

	for out, d := range data {
		t, err := template.ParseFiles("templates/enumerable.go.tpl")
		if err != nil {
			log.Println(err)
		}

		fp, err := os.Create(out)
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		d.Extras = []string{}

		for method, targets := range extrasMethods {
			for _, t := range targets {
				if t == d.Prefix {
					d.Extras = append(d.Extras, method)
					break
				}
			}
		}

		if err = t.Execute(fp, d); err != nil {
			log.Println(err)
		}
	}
}
