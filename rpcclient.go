package rpcclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client represents an RPC Client which helps interacting with either a Bitcoin or Litecoin RPC server.
// All of the functions will handle converting the returned types from the underlying JSON types.
type Client struct {
	// httpClient is the underlying HTTP client to use for the JSON-RPC requests.
	httpClient *http.Client

	// retryCount holds the number of times the client has tried to reconnect to the RPC server.
	retryCount int

	config *Config
}

// Config are the options for connecting to the RPC server.
type Config struct {
	// Host is the IP and port or FQDN of the RPC server you want to connect to.
	Host string

	// User is the username to use to authenticate to the RPC server.
	User string

	// Pass is the password to use to authenticate to the RPC server.
	Pass string

	// DisableTLS specifies whether TLS should be disabled. It is recommended to leave this enabled.
	DisableTLS bool

	// Certificates are the bytes for a PEM-encoded certificate chain used for the TLS connection.
	// It is ignored if DisableTLS is true.
	Certificates []byte

	// Proxy, ProxyUser and ProxyPass are for configuring the proxy http.Client should use.
	// ProxyUser and ProxyPass are ignored if Proxy is empty.
	Proxy     string
	ProxyUser string
	ProxyPass string

	// DisableAutoReconnect specifies whether the client should try to reconnect when the server has been disconnected.
	DisableAutoReconnect bool
}

// New creates a new *Client based on the provided config.
func New(config *Config) (*Client, error) {
	httpClient, err := newHTTPClient(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:     config,
		httpClient: httpClient,
	}

	return client, nil
}

// Request is a request to the JSON RPC server.
type Request struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// Response represents a response from the RPC server.
type Response struct {
	Result interface{} `json:"result"`
	Error  *RPCError   `json:"error"`
}

// RPCError represents an error returned by the RPC Server.
// It is included in the Response.
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendReq sends an HTTP POST request to the RPC server.
func (c *Client) SendReq(method string, result any, params ...any) error {
	rawReq := &Request{
		Method: method,
		Params: make([]interface{}, 0, len(params)),
	}

	rawReq.Params = append(rawReq.Params, params...)

	reqBody, err := json.Marshal(rawReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.config.Host, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.config.User, c.config.Pass)

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	err2 := res.Body.Close()
	if err != nil || err2 != nil {
		return err
	}

	rawResp := &Response{
		Result: result,
	}

	err = json.Unmarshal(body, &rawResp)
	if err != nil {
		return err
	}

	if rawResp.Error != nil {
		return fmt.Errorf("rpc response: code %v: %#v", rawResp.Error.Code, rawResp.Error.Message)
	}

	return nil
}

// newHTTPClient creates a new http.Client that is configured with the proxy and TLS settings in the Config.
func newHTTPClient(config *Config) (*http.Client, error) {
	var proxyFunc func(*http.Request) (*url.URL, error)
	if config.Proxy != "" {
		proxyUrl, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, err
		}

		proxyFunc = http.ProxyURL(proxyUrl)
	}

	var tlsConfig *tls.Config
	if !config.DisableTLS {
		if len(config.Certificates) > 0 {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(config.Certificates)
			tlsConfig = &tls.Config{
				RootCAs: pool,
			}
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
		},
	}

	return client, nil
}
