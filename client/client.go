package client

import (
	//	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

var (
	errNoEndpoints = errors.New("client: No endpoints provided")
	errNoUsername  = errors.New("client: No username provided")
	errNoPassword  = errors.New("client: No password provided")
	//errNoTLSConfig = errors.New("client: No TLS Config provided")
)

// Config configures a Client.
type Config struct {
	// Endpoints is a list of API endpoints
	Endpoints []string
	// Username specifies the user for the authorization header
	Username string
	// Password is the password for the specified user to add as
	// an authorization header to the request.
	Password string
	// DialTimeout is the timeout for dialing a client connection
	//DialTimeout time.Duration
	// Client TLS credentials
	//TLS *tls.Config
}

// Client provides a nx-api client session supporting the REST
// and ins_api client types.
// TODO: Add the REST Client type
type Client struct {
	InsAPI    insAPI `json:"ins_api"`
	netClient http.Client
}

// New creates a new Client from the given Config.
func New(config *Config) (*Client, error) {
	if len(config.Endpoints) == 0 {
		return nil, errNoEndpoints
	}
	if len(config.Username) == 0 {
		return nil, errNoUsername
	}
	if len(config.Password) == 0 {
		return nil, errNoPassword
	}
	return newClient(config), nil
}

// TODO: Add the REST Client
func newClient(config *Config) *Client {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	client := &Client{
		InsAPI: insAPI{
			Version:      Verson,
			Type:         CliType,
			Chunk:        Chunk,
			Sid:          Sid,
			Input:        Input,
			OutputFormat: OutputFormat,
		},
		netClient: *netClient,
	}
	return client
}
