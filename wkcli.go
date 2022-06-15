package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	refs = flag.Bool("l", false, "Return references from page")
	ifl  = flag.Bool("x", false, "Return exact match")
	wiki = flag.String("w", "en.wikipedia.org", "Mediawiki to search")
)

func main() {
	flag.Parse()
	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(1)
	}
	query := strings.Join(flag.Args(), "+")
	d, err := getInitial(query)
	if err != nil {
		log.Fatalf("Error in initial query: %v", err)
	}
	if isAmbiguous(d) && ! *ifl {
		// No unique result, print list and exit
		if err := listAmbiguities(d); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	d2, err := getLinks(query)
	defer d2.Close()
	if err != nil {
		log.Fatalf("%s\n%s\n", listHeading(d), err)
	}
	if err := listLinks(d, d2); err != nil {
		log.Fatal(err)
	}
}
