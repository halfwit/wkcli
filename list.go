package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

type fromInterface []string

func listAmbiguities(r []interface{}) error {
	var iface fromInterface
	if len(r) < 3 {
		return fmt.Errorf("Received incomplete results, cannot continue")
	}
	names := iface.getslice(r[1], "Names")
	syn   := iface.getslice(r[3], "Synopsis")
	if err := iface.errors(); err != nil {
		return err
	}
	for n, m := range names {
		fmt.Printf("%s - %s\n", m, syn[n])
	}
	return nil
}

func listLinks(header []interface{}, results io.ReadCloser) error {
	err := listHeading(header)
	if err != nil || ! *refs {
		return err
	}
	json, _ := ioutil.ReadAll(results)
	links := gjson.GetManyBytes(json, "*.extlinks.url")
	fmt.Println(links)
	return err
}

func listHeading(r []interface{}) error {
	var iface fromInterface
	if len(r) < 3 || r == nil {
		return fmt.Errorf("Unable to parse results, cannot continue")
	}
	name := iface.getslice(r[1], "Names")
	syns := iface.getslice(r[2], "Synopsis")
	url  := iface.getslice(r[3], "Urls")
	if err := iface.errors(); err != nil {
		return err
	}
	fmt.Printf("%s - %s\n%s\n", name[0], url[0], syns[0])
	return nil
}

func (e *fromInterface) getslice(i interface{}, name string) []interface{} {
	switch v := i.(type) {
	case []interface{}:
		return v
	}
	*e = append(*e, fmt.Sprintf("Unable to locate %s in results", name))
	return nil
}

func (e *fromInterface) errors() error {
	if len(*e) > 0 {
		return fmt.Errorf("%s\n", *e)
	}
	return nil
}
