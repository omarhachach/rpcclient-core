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

	"github.com/omarhachach/rpcclient-core/types"
)

// Enforce Client has to be implementation of IClient.
var _ IClient = &Client{}

// IClient is the interface representation of an rpcclient to a Bitcoin or Litecoin RPC server.
// This should be used when passing the client to methods or structs, to make the functions mockable and testable.
type IClient interface {
	// GetBlock returns hex-encoded block data.
	GetBlock(hash string) (string, error)
	// GetBlockVerbose returns a decoded block.
	GetBlockVerbose(hash string) (*types.Block, error)
	// GetBlockVerboseTx returns a decoded block with decoded transaction info.
	GetBlockVerboseTx(hash string) (*types.BlockTx, error)
	// GetBlockHash returns the hash of the block in the best-block-chain at the provided height
	GetBlockHash(height int) (string, error)
	// GetBlockHeader returns a hex-encoded block header.
	GetBlockHeader(hash string) (string, error)
	// GetBlockHeaderVerbose returns a decoded block header.
	GetBlockHeaderVerbose(hash string) (*types.BlockHeader, error)
	// GetBlockStats returns the computed per block stats for a given block.
	GetBlockStats(hash string) (*types.BlockStats, error)
	// GetBlockStatsHeight is the same GetBlockStats, but uses the block height instead of the hash.
	GetBlockStatsHeight(height int) (*types.BlockStats, error)
	// PreciousBlock treats a block as if it were received before others with the same work.
	PreciousBlock(hash string) error

	// GetBestBlockHash returns the hash of the best (tip) block in the most-work fully-validated chain.
	GetBestBlockHash() (string, error)
	// GetBlockChainInfo returns various state info regarding blockchain processing.
	GetBlockChainInfo() (*types.BlockChainInfo, error)
	// GetBlockCount returns the number of blocks in the longest chain.
	GetBlockCount() (int64, error)
	// GetBlockFilter retrieves the BIP 157 content filter for a particular block.
	GetBlockFilter(blockhash, filtertype string) (*types.BlockFilter, error)
	// GetChainTips returns information about all know tips in the block tree, including the main chain as well as
	// orphaned branches.
	GetChainTips() ([]*types.ChainTip, error)
	// GetChainTxStats computes statistics about the total number- and rate of transactions in the chain.
	GetChainTxStats(nblocks int, blockhash string) (*types.ChainTxStats, error)
	// GetDifficulty returns the proof-of-work difficulty as a multiple of the minimum difficulty
	GetDifficulty() (float64, error)
	// PruneBlockchain prunes the blockchain up to the given height.
	PruneBlockchain(height int) (int, error)
	// VerifyChain verifies the blockchain database.
	VerifyChain(level int) (bool, error)
	// GetMemoryInfo returns an object containing information about memory usage.
	GetMemoryInfo() (*types.MemoryInfo, error)
	// GetMemoryInfoMalloc returns an XML string describing low-level heap state. (Only available if node is compiled
	// with glibc 2.10+)
	GetMemoryInfoMalloc() (string, error)
	// GetRPCInfo returns details about the RPC server.
	GetRPCInfo() (*types.RPCInfo, error)

	// GenerateBlock mines a block with a set of ordered transactions immediately to a specified address or descriptor.
	// The txs param is either a raw transaction or a txid in the mempool.
	GenerateBlock(output string, txs []string) (*types.GenerateBlockResult, error)
	// GenerateToAddress mines blocks to a specified address. If maxtries is <=0, will be set to default of 1000000.
	GenerateToAddress(nblocks int, adress string, maxtries int) ([]string, error)
	// GenerateToDescriptor is the same as GenerateToAddress, except it uses a descriptor instead of an address.
	GenerateToDescriptor(nblocks int, descriptor string, maxtries int) ([]string, error)

	// GetMempoolAncestors gets a list of transaction ids for the in-mempool ancestors of the provided txid.
	GetMempoolAncestors(txid string) ([]string, error)
	// GetMempoolAncestorsVerbose is like GetMempoolAncestors but will map the transaction ids to detail objects.
	GetMempoolAncestorsVerbose(txid string) (map[string]*types.MempoolTransaction, error)
	// GetMempoolDescendants gets a list of transaction ids for the in-mempool descendants of the provided txid.
	GetMempoolDescendants(txid string) ([]string, error)
	// GetMempoolDescendantsVerbose is like GetMempoolDescendants but will map the transaction ids to detail objects.
	GetMempoolDescendantsVerbose(txid string) (map[string]*types.MempoolTransaction, error)
	// GetMempoolEntry retrieves the mempool data for a given transaction. (Txid must be in mempool).
	GetMempoolEntry(txid string) (*types.MempoolTransaction, error)
	// GetMempoolInfo returns details on the active state of the transaction memory pool.
	GetMempoolInfo() (*types.MempoolInfo, error)
	// GetRawMempool returns a list of txids in the mempool.
	GetRawMempool() ([]string, error)
	// GetRawMempoolVerbose is like GetRawMempool but will map the txids to detail objects.
	GetRawMempoolVerbose() (map[string]*types.MempoolTransaction, error)
	// SaveMempool dumps the mempool to disk.
	SaveMempool() error
	// Uptime returns the total uptime of the server in seconds.
	Uptime() (int, error)
	// Stop requests a graceful stop of the node.
	Stop() error

	// GetBlockTemplate returns data needed to construct a block to work on. If template is nil, will use default.
	GetBlockTemplate(template *types.BlockTemplateRequest) (*types.BlockTemplate, error)
	// GetMiningInfo returns mining-related information.
	GetMiningInfo() (*types.MiningInfo, error)
	// GetNetworkHashPS returns the estimated network hashes per second based on the latest nblocks.
	GetNetworkHashPS(nblocks, height int) (int, error)
	// PrioritiseTransaction accepts the transaction into mined blocks at a higher (or lower) priority.
	PrioritiseTransaction(txid string, feeDelate int) (bool, error)
	// SubmitBlock submits a new block to the network.
	SubmitBlock(hexdata string) error
	// SubmitHeader decodes the hexdata as a header and submits it as a candidate chain tip if valid.
	SubmitHeader(hexdata string) error

	// GetTxOut returns details about an unspent transaction output.
	GetTxOut(txid string, vout int, includeMempool bool) (*types.TransactionOut, error)
	// GetTxOutProof returns a hex-encoded proof that the transaction was included in a block. Read RPC docs for a note
	// on reliability.
	GetTxOutProof(txidsFilter []string) (string, error)
	// GetTxOutProofInBlock returns a hex-encoded proof that the transaction was included in the block. Read RPC docs
	// for a note on reliability.
	GetTxOutProofInBlock(txidsFilter []string, blockhash string) (string, error)
	// GetTxOutSetInfo returns statistics about the unspect transaction output set.
	GetTxOutSetInfo() (*types.TransactionOutSetInfo, error)
	// ScanTxOutSet is experimental. Please read the docs https://developer.bitcoin.org/reference/rpc/scantxoutset.html.
	ScanTxOutSet(action types.ScanTxOutSetObject, scanObjects ...*types.ScanTxOutSetObject) (*types.ScanTxOutSetDetails, error)
	// VerifyTxOutProof verifies the proof points to a transaction in a block.
	VerifyTxOutProof(proof string) ([]string, error)
	// AnalyzePSBT analyzes and provides information about the current status of a AnalyzePSBTResult and its inputs.
	AnalyzePSBT(psbtbase64 string) (*types.AnalyzePSBTResult, error)
	// CombinePSBT combines multiple PSBTs into one.
	CombinePSBT(psbts []string) (string, error)
	// CombineRawTransaction combines multiple partially signed transaction into one transaction.
	CombineRawTransaction(txs []string) (string, error)
	// ConvertToPSBT converts a transaction to an psbt. If iswitness is null, it will use a heuristic to determine it.
	ConvertToPSBT(hex string, permitsigdata bool, iswitness *bool) (string, error)
	// CreatePSBT creates a psbt.
	CreatePSBT(inputs []*types.CreateTxInput, outputs []map[string]string, locktime int, replaceable bool) (string, error)
	// CreateRawTransaction creates a raw transaction.
	CreateRawTransaction(inputs []*types.CreateTxInput, outputs []map[string]string, locktime int, replacable bool) (string, error)
	// DecodePSBT takes a base64 psbt string and converts it to an object.
	DecodePSBT(psbtbase64 string) (*types.PSBT, error)
	// DecodeRawTransaction takes a hex transactiopn and converts it to an object. If is witness is null, it will use a
	// heuristic to determine it.
	DecodeRawTransaction(txhex string, iswitness *bool) (*types.Transaction, error)
	// DecodeScript decodes a hex-encoded script.
	DecodeScript(scripthex string) (*types.DecodedScript, error)
	// FinalizePSBT finalizes the inputs of a PSBT.
	FinalizePSBT(psbtbase64 string, extract bool) (*types.FinalizePSBTResult, error)
	// FundRawTransaction will selects inputs to meet its outputs value. If iswitness is null, it will use a heuristic
	// to determine it.
	FundRawTransaction(tx string, opts *types.FundRawTransactionOptions, iswitness *bool) (*types.FundRawTransactionResult, error)
	// GetRawTransaction get a transaction from mempool or the blockchain. If blockhash is not nil, will use the
	// blockhash to look for the transaction.
	GetRawTransaction(txid string, blockhash *string) (string, error)
	// GetRawTransactionVerbose gets a transaciton from mempool or the blockchain. If blockhash is not nil, will use the
	// blockhash to look for the transaction.
	GetRawTransactionVerbose(txid string, blockhash *string) (*types.Transaction, error)
	// JoinPSBTs joins multiple distinct PSBTs with different inputs and outputs into one PSBT.
	JoinPSBTs(psbts []string) (string, error)
	// SendRawTransaction sends a transaction the local node and network. If maxfeerate is nil, it will use node default.
	SendRawTransaction(hex string, maxfeerate *float64) (string, error)
	// SignRawTransactionWithKey signs a raw transaction with the provided keys. If prevTxs is null or length 0, it will
	// be omitted. If sigHashType is "" will be set to types.SigHashTypeAll.
	SignRawTransactionWithKey(hex string, privKeys []string, prevTxs []*types.PreviousTransaction, sigHashTypes types.SigHashType) (*types.SignRawTransactionResult, error)
	// TestMempoolAccept returns the result of mempool acceptance tsts indicating if raw transaction would be accepted by
	// the mempool.
	TestMempoolAccept(rawtxs []string, maxfeeRate *float64) ([]*types.TestMempoolAcceptResult, error)
	// UtxoUpdatePSBT updates all segwit inputs and outputs in a PSBT with data from output descriptors, the UTXO set or the
	// mempool.
	UtxoUpdatePSBT(psbt string, scanObjects ...*types.ScanTxOutSetObject) (string, error)

	// EstimateSmartFee estimates the approximate fee per kilobyte needed for a transaction for a transaction to begin
	// within confTarget blocks. If estimateMode is nil, will use default.
	EstimateSmartFee(confTarget int, estimateMode *types.EstimateMode) (*types.EstimateSmartFeeResult, error)
}

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

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
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
