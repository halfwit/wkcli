package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var initial = "https://%s/w/api.php?action=opensearch&search=%s&namespace=0&format=json"
var links   = "https://%s/w/api.php?action=query&titles=%s&prop=extlinks&ellimit=500&formatversion=2&format=json"

func getInitial(query string) ([]interface{}, error) {
	url := fmt.Sprintf(initial, *wiki, query)
	var r []interface{}
	gr, err := run(url)
	defer gr.Close()
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(gr).Decode(&r); err != nil {
		return nil, err
	}
	return r, nil
}

func getLinks(query string) (io.ReadCloser, error) {
	url := fmt.Sprintf(links, *wiki, query)
	return run(url)	
}

func isAmbiguous(r []interface{}) bool {
	// The codepath for true will immediately catch and properly print this error
	if len(r) < 1 {
		return true
	}
	names, ok := r[1].([]interface{})
	if ! ok {
		return true
	}
	return len(names) > 1
}

func run(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "wkcli")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}