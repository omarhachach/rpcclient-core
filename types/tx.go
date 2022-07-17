package types

import (
	"encoding/json"
	"errors"
)

// Transaction represents a transaction in a block.
type Transaction struct {
	// InActiveChain whether the block is in the active chain or not (only present with explicit "blockhash" argument).
	InActiveChain bool `json:"inActiveChain"`
	// Hex is the hex-encoded data for Txid.
	Hex string `json:"hex"`
	// Txid is the transaction id.
	Txid string `json:"txid"`
	// Hash is the transaction hash (differs from Txid for witness transactions).
	Hash string `json:"hash"`
	// Size is the serialized transaction size.
	Size int `json:"size"`
	// Vsize is the virtual transaction size (differs from Size for witness transactions).
	Vsize int `json:"vsize"`
	// Weight is the transaction's weight (between Vsize*4-3 and Vsize*4)
	Weight int `json:"weight"`
	// Version is the version.
	Version int `json:"version"`
	// Locktime is the locktime.
	Locktime int `json:"locktime"`
	// Vin are the transaction inputs.
	Vin []*Vin `json:"vin"`
	// Vout are the transaction outputs.
	Vout []*Vout `json:"vout"`
	// Blockhash is the hash of the block where this transaction is included.
	Blockhash string `json:"blockhash"`
	// Confirmations are the number of confirmations. (Same as Block.Confirmations for the Blockhash)
	Confirmations int `json:"confirmations"`
	// Blocktime is the block time expressed in UNIX epoch time.
	Blocktime int `json:"blocktime"`
	// Time is the same as Blocktime.
	Time int `json:"time"`
}

// Vin represents a transaction input.
type Vin struct {
	// Txid is the transaction id of the input.
	Txid string `json:"txid"`
	// Vout is the output number from the transaction found with Txid.
	Vout int `json:"vout"`
	// ScriptSig is the script.
	ScriptSig *ScriptSig `json:"scriptSig"`
	// Sequence is the script sequence number.
	Sequence int `json:"sequence"`
	// Txinwitness is a list of hex-encoded witness data.
	Txinwitness []string `json:"txinwitness"`
}

// ScriptSig represents a script for an input.
type ScriptSig struct {
	// Asm is the script in asm.
	Asm string `json:"asm"`
	// Hex is the script in hex.
	Hex string `json:"hex"`
}

// Vout represents a transaction output.
type Vout struct {
	// Value is the value of the ouput in BTC.
	Value float64 `json:"value"`
	// N is the index of the output.
	N int `json:"n"`
	// ScriptPubKey is the output script.
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey"`
}

// RedeemScript contains info about a redeem script.
type RedeemScript struct {
	*ScriptSig
	// Type is the output type, eg 'pubkeyhash'.
	Type string `json:"type"`
}

// ScriptPubKey is the output script.
type ScriptPubKey struct {
	*RedeemScript
	// ReqSigs is the required amount signatures.
	ReqSigs int `json:"reqSigs"`
	// Address is the output address (for bitcoin-core >22)
	Address string `json:"address"`
	// Addresses are the output addresses (for bitcoin-core and litecoin-core).
	Addresses []string `json:"addresses"`
}

// GetAddress retrieves the first address, it first checks Address then first entry of Addresses.
// Returns "" if none is found.
func (s *ScriptPubKey) GetAddress() string {
	if s.Address != "" {
		return s.Address
	}

	if len(s.Addresses) > 0 {
		return s.Addresses[0]
	}

	return ""
}

// TransactionOut holds info about a transaction's unspent output.
type TransactionOut struct {
	BestBlock     string        `json:"bestblock"`
	Confirmations int           `json:"confirmations"`
	Value         float64       `json:"value"`
	ScriptPubKey  *ScriptPubKey `json:"scriptPubKey"`
	Coinbase      bool          `json:"coinbase"`
}

// TransactionOutSetInfo holds stats about an unspent transaciton output set.
type TransactionOutSetInfo struct {
	Height          int     `json:"height"`
	BestBlock       string  `json:"bestblock"`
	Transactions    int     `json:"transactions"`
	TxOuts          int     `json:"txouts"`
	BogoSize        int     `json:"bogosize"`
	HashSerialized2 string  `json:"hash_serialized_2"`
	DiskSize        int     `json:"disk_size"`
	TotalAmount     float64 `json:"total_amount"`
}

// ScanTxOutSetAction is an enum with actions for ScanTxOutSet.
type ScanTxOutSetAction string

const (
	ScanTxOutSetStart  ScanTxOutSetAction = "start"
	ScanTxOutSetAbort  ScanTxOutSetAction = "abort"
	ScanTxOutSetStatus ScanTxOutSetAction = "status"
)

// ScanTxOutSetObject comes with the special ToJSON method.
// This method will be called when sent to the RPC endpoint. This is since the argument is either an object or string.
// If Descriptor is set, it will return that regardless of everything else.
// Otherwise it will send the object with Desc and optionally RangeN or Range.
type ScanTxOutSetObject struct {
	Descriptor string
	Desc       string
	Range      int
	RangeN     []int
}

type scanTxOutObjectInt struct {
	Desc  string `json:"desc"`
	Range int    `json:"range,omitempty"`
}

type scanTxOutObjectArr struct {
	Desc  string `json:"desc"`
	Range []int  `json:"range,omitempty"`
}

// ToJSON serializes the object according to https://developer.bitcoin.org/reference/rpc/scantxoutset.html.
func (s *ScanTxOutSetObject) ToJSON() ([]byte, error) {
	if s.Descriptor != "" {
		return []byte(s.Descriptor), nil
	}

	if s.Desc == "" {
		return nil, errors.New("ScanTxOutSetObject can not have both an empty Desc and Descriptor")
	}

	if len(s.RangeN) == 2 {
		return json.Marshal(&scanTxOutObjectArr{
			Desc:  s.Desc,
			Range: s.RangeN,
		})
	}

	return json.Marshal(&scanTxOutObjectInt{
		Desc:  s.Desc,
		Range: s.Range,
	})
}

// ScanTxOutSetDetails is the result of the ScanTxOutSet request.
type ScanTxOutSetDetails struct {
	Success     bool                   `json:"success"`
	TxOuts      int                    `json:"txouts"`
	Height      int                    `json:"height"`
	BestBlock   string                 `json:"bestblock"`
	Unspents    []*ScanTxOutSetUnspent `json:"unspents"`
	TotalAmount float64                `json:"total_amount"`
}

// ScanTxOutSetUnspent reprents and unspent output returned in ScanTxOutSetDetails.
type ScanTxOutSetUnspent struct {
	Txid         string  `json:"txid"`
	Vout         int     `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	Desc         string  `json:"desc"`
	Amount       float64 `json:"amount"`
	Height       int     `json:"height"`
}

// AnalyzePSBTResult is contains details of a PSBT analysis.
type AnalyzePSBTResult struct {
	Inputs          []*AnalyzePSBTInput `json:"inputs"`
	EstimatedVsize  int                 `json:"estimated_vsize,omitempty"`
	EstimateFeerate float64             `json:"estimate_feerate,omitempty"`
	Fee             float64             `json:"fee,omitempty"`
	Next            string              `json:"next"`
	Error           string              `json:"error,omitempty"`
}

// AnalyzePSBTInput is the input of a AnalyzePSBTResult.
type AnalyzePSBTInput struct {
	HasUTXO  bool              `json:"has_utxo"`
	IsFinal  bool              `json:"is_final"`
	Missings *PSBTInputMissing `json:"missings,omitempty"`
	Next     string            `json:"next,omitempty"`
}

// PSBTInputMissing holds the things that are missing from a AnalyzePSBTInput.
type PSBTInputMissing struct {
	Pubkeys       []string `json:"pubkeys,omitempty"`
	Signatures    []string `json:"signatures,omitempty"`
	RedeemScript  string   `json:"redeemscript,omitempty"`
	WitnessScript string   `json:"witnessscript,omitempty"`
}

// CreateTxInput is an input for creating a raw transaction or psbt.
type CreateTxInput struct {
	Txid     string `json:"txid"`
	Vout     int    `json:"vout"`
	Sequence int    `json:"sequence,omitempty"`
}

// PSBT represents a partially signed transaction.
type PSBT struct {
	Tx      *Transaction      `json:"tx"`
	Unknown map[string]string `json:"unknown"`
	Inputs  []*PSBTInput      `json:"inputs"`
	Outputs []*PSBTOutput     `json:"outputs"`
	Fee     float64           `json:"fee,omitempty"`
}

// PSBTInput contains information about a PSBT's input.
type PSBTInput struct {
	NonWitnessUTXO     *Transaction      `json:"non_witness_utxo,omitempty"`
	WitnessUTXO        *PSBTWitnessUTXO  `json:"witness_utxo,omitempty"`
	PartialSignatures  map[string]string `json:"partial_signatures,omitempty"`
	Sighash            string            `json:"sighash,omitempty"`
	RedeemScript       *RedeemScript     `json:"redeem_script,omitempty"`
	WitnessScript      *RedeemScript     `json:"witness_script,omitempty"`
	Bip32Derivs        []*Bip32Deriv     `json:"bip32_derivs,omitempty"`
	FinalScriptSig     *ScriptSig        `json:"final_scriptsig,omitempty"`
	FinalScriptwitness []string          `json:"final_scriptwitness"`
	Unknown            map[string]string `json:"unknown"`
}

// PSBTWitnessUTXO is a transaction output for witness utxo.
type PSBTWitnessUTXO struct {
	Amount       float64       `json:"amount"`
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey"`
}

// Bip32Deriv is a public key with the derivation path.
type Bip32Deriv struct {
	MasterFingerprint string `json:"master_fingerprint"`
	Path              string `json:"path"`
}

// PSBTOutput contains information about a PSBT's output.
type PSBTOutput struct {
	RedeemScript  *RedeemScript     `json:"redeem_script,omitempty"`
	WitnessScript *RedeemScript     `json:"witness_script,omitempty"`
	Bip32Derivs   []*Bip32Deriv     `json:"bip32_derivs,omitempty"`
	Unknown       map[string]string `json:"unknown"`
}

// DecodedScript represents a decoded hex-encoded script.
type DecodedScript struct {
	Asm       string        `json:"asm"`
	Type      string        `json:"type"`
	ReqSigs   int           `json:"reqSigs"`
	Addresses []string      `json:"addresses"`
	P2SH      string        `json:"p2sh"`
	Segwit    *SegwitScript `json:"segwit"`
}

// SegwitScript is the result of a witness script public key.
type SegwitScript struct {
	Asm        string   `json:"asm"`
	Hex        string   `json:"hex"`
	Type       string   `json:"type"`
	ReqSigs    int      `json:"reqSigs"`
	Addresses  []string `json:"addresses"`
	P2SHSegwit string   `json:"p2sh-segwit"`
}

// FinalizePSBTResult is the result of finalizepsbt.
type FinalizePSBTResult struct {
	PSBT     string `json:"psbt"`
	Hex      string `json:"hex"`
	Complete bool   `json:"complete"`
}

// EstimateMode represents the fee estimate mode.
type EstimateMode string

const (
	EstimateModeConservative EstimateMode = "conservative"
	EstimateModeUnset        EstimateMode = "unset"
	EstimateModeEconomical   EstimateMode = "economical"
)

// FundRawTransactionOptions contains options for fundrawtransaction.
type FundRawTransactionOptions struct {
	AddInputs              bool         `json:"add_inputs,omitempty"`
	ChangeAddress          string       `json:"changeAddress,omitempty"`
	ChangePosition         *int         `json:"changePosition,omitempty"`
	ChangeType             string       `json:"change_type,omitempty"`
	IncludeWatching        bool         `json:"includeWatching"`
	LockUnspents           bool         `json:"lockUnspents"`
	FeeRateSat             int          `json:"fee_rate,omitempty"`
	FeeRate                float64      `json:"feeRate,omitempty"`
	SubtractFeeFromOutputs []int        `json:"subtractFeeFromOutputs,omitempty"`
	Replaceable            bool         `json:"replaceable"`
	ConfTarget             int          `json:"conf_target,omitempty"`
	EstimateMode           EstimateMode `json:"estimate_mode,omitempty"`
}

// FundRawTransactionResult is the result of fundrawtransaction.
type FundRawTransactionResult struct {
	Hex       string  `json:"hex"`
	Fee       float64 `json:"fee"`
	ChangePos int     `json:"changepos"`
}

// PreviousTransaction represents an UTXO.
type PreviousTransaction struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	RedeemScript  string  `json:"redeemScript"`
	WitnessScript string  `json:"witness_script"`
	Amount        float64 `json:"amount"`
}

// SigHashType indicates a signature hash's type.
type SigHashType string

// The valid values for the SigHashType enum.
const (
	SigHashTypeAll                SigHashType = "ALL"
	SigHashTypeNone               SigHashType = "NONE"
	SigHashTypeSingle             SigHashType = "SINGLE"
	SigHashTypeAllAnyoneCanPay    SigHashType = "ALL|ANYONECANPAY"
	SigHashTypeNoneAnyoneCanPay   SigHashType = "NONE|ANYONECANPAY"
	SigHashTypeSingleAnyoneCanPay SigHashType = "SINGLE|ANYONECANPAY"
)

// SignRawTransactionResult is the result of signrawtransactionwithkey.
type SignRawTransactionResult struct {
	Hex      string                           `json:"hex"`
	Complete bool                             `json:"complete"`
	Errors   []*SignRawTransactionResultError `json:"errors"`
}

// SignRawTransactionResultError is included if there is an error when calling signrawtranscationwithkey.
type SignRawTransactionResultError struct {
	Txid      string `json:"txid"`
	Vout      int    `json:"vout"`
	ScriptSig string `json:"scriptSig"`
	Sequence  int    `json:"sequence"`
	Error     string `json:"error"`
}

// TestMempoolAcceptResult is the result of testmempoolaccept.
type TestMempoolAcceptResult struct {
	Txid    string `json:"txid"`
	Allowed bool   `json:"allowed"`
	Vsize   int
	Fees    struct {
		Base float64 `json:"base"`
	} `json:"fees"`
	RejectReason string `json:"reject-reason"`
}
