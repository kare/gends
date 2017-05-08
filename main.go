package main // import "kkn.fi/cmd/gends"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type ds struct {
	ShortTypeName string
	TypeName      string
	Type          string
	TypeZeroValue string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gends: ")
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: gends templates.json datastructure.tmpl")
		os.Exit(1)
	}

	tmpl, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var templates []ds
	err = json.Unmarshal(tmpl, &templates)
	if err != nil {
		log.Fatal(err)
	}

	json, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("default").Parse(string(json))
	if err != nil {
		log.Fatal(err)
	}
	for _, template := range templates {
		var (
			name             = fmt.Sprintf("%v.go", template.Type)
			mode             = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			perm os.FileMode = 0644
		)
		file, err := os.OpenFile(name, mode, perm)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if err := t.Execute(file, template); err != nil {
			log.Fatal(err)
		}
	}
}
