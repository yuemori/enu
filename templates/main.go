package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

func main() {
	flag.Parse()
	in := flag.Arg(0)
	out := flag.Arg(1)

	data := struct {
		Type               string
		TypeWithConstraint string
		Prefix             string
		ItemType           string
		ImportPkg          string
	}{
		Type:               flag.Arg(2),
		TypeWithConstraint: flag.Arg(3),
		Prefix:             flag.Arg(4),
		ItemType:           flag.Arg(5),
		ImportPkg:          flag.Arg(6),
	}

	t, err := template.ParseFiles(in)
	if err != nil {
		log.Println(err)
	}

	fp, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	if err = t.Execute(fp, data); err != nil {
		log.Println(err)
	}
}
