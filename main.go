package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
)

var (
	// Defaults to info logging.
	log = logrus.New()
)

// Client represents the nx-api ins_api endpoint request.
type Client struct {
	API InsAPI `json:"ins_api"`
}

// InsAPI represents an ins_api request.
type InsAPI struct {
	Version      string `json:"version"`
	Type         string `json:"type"`
	Chunk        string `json:"chunk"`
	Sid          string `json:"sid"`
	Input        string `json:"input"`
	OutputFormat string `json:"output_format"`
}

// Flags are command line flags passed to the client.
type Flags struct {
	address  string
	username string
	password string
	logLevel string
}

func main() {
	var flags Flags
	flag.StringVar(&flags.address, "address", "10.10.10.254", "IP address of nx-api endpoint")
	flag.StringVar(&flags.username, "username", "jdoe", "username of nx-api endpoint")
	flag.StringVar(&flags.password, "password", "changeme", "password of nx-api endpoint")
	// Log levels https://github.com/Sirupsen/logrus/blob/master/logrus.go#L36
	flag.StringVar(&flags.logLevel, "log-level", "info", "Set the logging level")
	// parse command-line flags
	flag.Parse()
	// validate arguments
	if url, err := url.Parse(flags.address); err != nil || url.String() == "" {
		log.Fatal("A valid IP address of nx-api endpoint is required")
	}
	// logging setup
	lvl, err := logrus.ParseLevel(flags.logLevel)
	if err != nil {
		log.Fatalf("invalid log-level: %v", err)
	}
	log.Level = lvl
	// issue a request to the nx-api
	resp, err := insAPIRequest(flags.username, flags.password, flags.address)
	if err != nil {
		log.Fatalf("Could not issue request to nx-api: %v", err)
	} else {
		// Note that this will require you add fmt to your list of imports.
		fmt.Println(string(resp))
	}
}

func insAPIRequest(username, password, address string) ([]byte, error) {
	// Initialiaze an InsAPIClient
	c := Client{
		API: InsAPI{
			Version:      "1.2",
			Type:         "cli_show",
			Chunk:        "0",
			Sid:          "1",
			Input:        "show version",
			OutputFormat: "json",
		},
	}
	// Marshall the InsAPIClient client.
	cbuf, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Build the request
	req, err := http.NewRequest("POST", "http://"+address+"/ins", bytes.NewBuffer(cbuf))
	if err != nil {
		return nil, err
	}
	// The header has to be set to application/json
	req.Header.Set("content-type", "application/json")
	// Set the credentials in the header.
	req.SetBasicAuth(username, password)
	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Defer the closing of the body
	defer resp.Body.Close()
	// Read the content into an array of bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Return the bytes
	return body, nil
}
