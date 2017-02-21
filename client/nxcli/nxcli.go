package client

import (
	//	"encoding/json"
	"errors"
	//	"fmt"
	//	"io"
	//	"math/rand"
	//	"net"
	"net/http"
	//	"net/url"
	//	"sort"
	//  "strconv"
	//	"sync"
	"time"
)

var (
	ErrNoEndpoints        = errors.New("client: no endpoints available")
	DefaultRequestTimeout = 5 * time.Second
)

// HTTP methods and version supported by the ins_api API
const (
	InsAPI       = "ins_api"
	Verson       = "1.2"
	CliType      = "cli_show"
	Chunk        = "0"
	Sid          = "1"
	Input        = "show version"
	OutputFormat = "json"
	Put          = "PUT"
)

type Request *http.Request
type Response *http.Response

//type cliType string

/*func(c cliType) cliShow() string {
	return "cli_show"
}*/

/*func(c cliType) cliConf() string {
	return "cli_conf"
}*/

// InsClient represents the nx-api ins_api client.
//type InsClient struct {
//	insAPI string `json:"ins_api"`
//}

// insAPI represents the
type insAPI struct {
	Version      string  `json:"version"`
	Type         string  `json:"type"`
	Chunk        string  `json:"chunk,omitempty"`
	Sid          string  `json:"sid"`
	Input        string  `json:"input,omitempty"`
	OutputFormat string  `json:"output_format,omitempty"`
	Outputs      Outputs `json:"outputs,omitempty"`
}

type Outputs struct {
	Output []struct {
		Msg      string   `json:"msg"`
		Code     string   `json:"code"`
		Body     struct{} `json:"body,omitempty"`
		Clierror string   `json:"clierror,omitempty"`
	} `json:"output"`
}
