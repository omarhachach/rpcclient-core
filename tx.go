package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GetTxOut returns details about an unspent transaction output.
func (c *Client) GetTxOut(txid string, vout int, includeMempool bool) (*types.TransactionOut, error) {
	var txout *types.TransactionOut

	return txout, c.SendReq("gettxout", &txout, txid, vout, includeMempool)
}

// GetTxOutProof returns a hex-encoded proof that the transaction was included in a block.
// Read RPC docs for a note on reliability.
func (c *Client) GetTxOutProof(txidsFilter []string) (string, error) {
	var proof string

	return proof, c.SendReq("gettxoutproof", &proof, txidsFilter)
}

// GetTxOutProofInBlock returns a hex-encoded proof that the transaction was included in the block.
// Read RPC docs for a note on reliability.
func (c *Client) GetTxOutProofInBlock(txidsFilter []string, blockhash string) (string, error) {
	var proof string

	return proof, c.SendReq("gettxoutproof", &proof, txidsFilter, blockhash)
}

// GetTxOutSetInfo returns statistics about the unspect transactio output set.
func (c *Client) GetTxOutSetInfo() (*types.TransactionOutSetInfo, error) {
	var info *types.TransactionOutSetInfo

	return info, c.SendReq("gettxoutsetinfo", &info)
}

// ScanTxOutSet is experimental. Please read the docs https://developer.bitcoin.org/reference/rpc/scantxoutset.html.
func (c *Client) ScanTxOutSet(action types.ScanTxOutSetObject, scanObjects ...*types.ScanTxOutSetObject) (*types.ScanTxOutSetDetails, error) {
	var details *types.ScanTxOutSetDetails

	objs := make([]string, len(scanObjects))
	for idx, obj := range scanObjects {
		serializedObj, err := obj.ToJSON()
		if err != nil {
			return nil, err
		}

		objs[idx] = string(serializedObj)
	}

	return details, c.SendReq("scantxoutset", &details, action, objs)
}

// VerifyTxOutProof verifies that proof points to a transaction in ablock.
func (c *Client) VerifyTxOutProof(proof string) ([]string, error) {
	var txids []string

	return txids, c.SendReq("verifytxoutproof", &txids, proof)
}

// AnalyzePSBT analyzes and provides information about the current status of a AnalyzePSBTResult and its inputs.
func (c *Client) AnalyzePSBT(psbtbase64 string) (*types.AnalyzePSBTResult, error) {
	var psbt *types.AnalyzePSBTResult

	return psbt, c.SendReq("analyzepsbt", &psbt, psbtbase64)
}

// CombinePSBT combines multiple PSBTs into one.
func (c *Client) CombinePSBT(psbts []string) (string, error) {
	var psbt string

	return psbt, c.SendReq("combinepsbt", &psbt, psbts)
}

// CombineRawTransaction combines multiple partially signed transaction into one transaction.
func (c *Client) CombineRawTransaction(txs []string) (string, error) {
	var tx string

	return tx, c.SendReq("combinerawtransaction", &tx, txs)
}

// ConvertToPSBT converts a transaction to a AnalyzePSBTResult.
// If iswitness is null, it will use a heuristic to determine it.
func (c *Client) ConvertToPSBT(hex string, permitsigdata bool, iswitness *bool) (string, error) {
	var psbt string

	if iswitness != nil {
		return psbt, c.SendReq("converttopsbt", &psbt, hex, permitsigdata, *iswitness)
	}

	return psbt, c.SendReq("converttopsbt", &psbt, hex, permitsigdata)
}

// CreatePSBT creates a AnalyzePSBTResult.
func (c *Client) CreatePSBT(inputs []*types.CreateTxInput, outputs []map[string]string, locktime int, replaceable bool) (string, error) {
	var psbt string

	return psbt, c.SendReq("createpsbt", &psbt, inputs, outputs, locktime, replaceable)
}

// CreateRawTransaction creates a raw transaction.
func (c *Client) CreateRawTransaction(inputs []*types.CreateTxInput, outputs map[string]string, locktime int, replaceable bool) (string, error) {
	var rawtx string

	return rawtx, c.SendReq("createrawtransaction", &rawtx, inputs, outputs, locktime, replaceable)
}

// DecodePSBT takes a base64 psbt string and converts it to an object.
func (c *Client) DecodePSBT(psbtbase64 string) (*types.PSBT, error) {
	var psbt *types.PSBT

	return psbt, c.SendReq("decodepsbt", &psbt, psbtbase64)
}

// DecodeRawTransaction takes a hex transaction and converts it to an object.
// If iswitness is null it will use a heuristic to determine it.
func (c *Client) DecodeRawTransaction(txhex string, iswitness *bool) (*types.Transaction, error) {
	var tx *types.Transaction

	if iswitness != nil {
		return tx, c.SendReq("decoderawtransaction", &tx, txhex, *iswitness)
	}

	return tx, c.SendReq("decoderawtransaction", &tx, txhex)
}

// DecodeScript decodes a hex-encoded script.
func (c *Client) DecodeScript(scripthex string) (*types.DecodedScript, error) {
	var script *types.DecodedScript

	return script, c.SendReq("decodescript", &script, scripthex)
}

// FinalizePSBT finalizes the inputs of a PSBT.
func (c *Client) FinalizePSBT(psbtbase64 string, extract bool) (*types.FinalizePSBTResult, error) {
	var res *types.FinalizePSBTResult

	return res, c.SendReq("finalizepsbt", &res, psbtbase64, extract)
}

// FundRawTransaction will select inputs to meet its output value..
// If iswitness is null it will use a heuristic to determine it.
func (c *Client) FundRawTransaction(tx string, opts *types.FundRawTransactionOptions, iswitness *bool) (*types.FundRawTransactionResult, error) {
	var res *types.FundRawTransactionResult

	if iswitness != nil {
		return res, c.SendReq("fundrawtransaction", &res, tx, opts, *iswitness)
	}

	return res, c.SendReq("fundrawtransaction", &res, tx, opts)
}

// GetRawTransaction gets a transaction from mempool or the blockchain.
// If blockhash is not nil, will use the blockhash to look for the transaction.
func (c *Client) GetRawTransaction(txid string, blockhash *string) (string, error) {
	var tx string

	if blockhash != nil {
		return tx, c.SendReq("getrawtransaction", &tx, txid, false, *blockhash)
	}

	return tx, c.SendReq("getrawtransaction", &tx, txid, false)
}

// GetRawTransactionVerbose gets a transaction from mempool or the blockchain.
// If blockhash is not nil, will use the blockhash to look for the transaction.
func (c *Client) GetRawTransactionVerbose(txid string, blockhash *string) (*types.Transaction, error) {
	var tx *types.Transaction

	if blockhash != nil {
		return tx, c.SendReq("getrawtransaction", &tx, txid, true, *blockhash)
	}

	return tx, c.SendReq("getrawtransaction", &tx, txid, true)
}

// JoinPSBTs joins multiple distinct PSBTs with different inputs and outputs into one PSBT.
func (c *Client) JoinPSBTs(psbts []string) (string, error) {
	var psbt string

	return psbt, c.SendReq("joinpsbts", &psbt, psbts)
}

// SendRawTransaction sends a transaction to the local node and network.
// If maxfeerate is nil it will use node default.
func (c *Client) SendRawTransaction(hex string, maxfeerate *float64) (string, error) {
	var tx string

	if maxfeerate != nil {
		return tx, c.SendReq("sendrawtransaction", &tx, hex, *maxfeerate)
	}

	return tx, c.SendReq("sendrawtransaction", &tx, hex)
}

// SignRawTransactionWithKey signs a raw transaction witht he provided keys.
// If prevTxs is null or length 0, will be omitted. If sigHashType is "" will be set to types.SigHashTypeAll.
func (c *Client) SignRawTransactionWithKey(hex string, privKeys []string, prevTxs []*types.PreviousTransaction, sigHashType types.SigHashType) (*types.SignRawTransactionResult, error) {
	var res *types.SignRawTransactionResult

	if prevTxs == nil || len(prevTxs) == 0 {
		return res, c.SendReq("signrawtransactionwithkey", &res, hex, privKeys)
	}

	if sigHashType == "" {
		sigHashType = types.SigHashTypeAll
	}

	return res, c.SendReq("signrawtransactionwithkey", &res, hex, privKeys, prevTxs, sigHashType)
}
