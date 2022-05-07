package types

// BlockHeader holds meta-data about a block.
type BlockHeader struct {
	// Hash is the block hash.
	Hash string `json:"hash"`
	// Confirmations is the number of confirmations, or -1 if the block isn't on the main chain.
	Confirmations int `json:"confirmations"`
	// Height is the block height or index.
	Height int `json:"height"`
	// Version is the block version.
	Version int `json:"version"`
	// Versionhex is the block version formatted in hexadecimal.
	VersionHex string `json:"versionHex"`
	// Merkleroot is the merkle root.
	Merkleroot string `json:"merkleroot"`
	// Time is the block time expressed in UNIX epoch time.
	Time int `json:"time"`
	// Mediantime is the median block time expressed in UNIX epoch time.
	Mediantime int `json:"mediantime"`
	// Nonce is the nonce.
	Nonce int `json:"nonce"`
	// Bits are the bits.
	Bits string `json:"bits"`
	// Difficulty is the block difficulty.
	Difficulty float64 `json:"difficulty"`
	// Chainwork is the expected number of hashes required to produce the chain up to this block (in hex).
	Chainwork string `json:"chainwork"`
	// NTx is the number of transactions included in the block.
	NTx int `json:"nTx"`
	// Previousblockhash is the hash of the previous block in the chain.
	Previousblockhash string `json:"previousblockhash"`
	// Nextblockhash is the hash of the next block in the chain.
	Nextblockhash string `json:"nextblockhash"`
}

// Block represents a block on the chain and contains information about it and the transactions included in this block.
type Block struct {
	*BlockHeader
	// Size is the block size in bytes.
	Size int `json:"size"`
	// StrippedSize is the block size excluding the witness data in bytes.
	StrippedSize int `json:"strippedSize"`
	// Weight is the block weight as defined in BIP 141.
	Weight int `json:"weight"`
	// Tx is a list of transaction ids included in this block.
	Tx []string `json:"tx"`
}

// BlockTx is identital to Block but includes verbose Transaction information.
type BlockTx struct {
	*BlockHeader
	// Size is the block size in bytes.
	Size int `json:"size"`
	// StrippedSize is the block size excluding the witness data in bytes.
	StrippedSize int `json:"strippedSize"`
	// Weight is the block weight as defined in BIP 141.
	Weight int `json:"weight"`
	// Tx is a list of transaction ids included in this block.
	Tx []*Transaction `json:"tx"`
}

// BlockStats holds calculated statistics about a block. All values are in satoshis.
type BlockStats struct {
	// AverageFee is the average fee in this block.
	AverageFee int `json:"avgfee"`
	// AverageFeeRate is the average fee rate (in satoshis per virtual byte, see Transaction.Vsize).
	AverageFeeRate int `json:"avgfeerate"`
	// AverageTxSize is the average transaction size.
	AverageTxSize int `json:"avgtxsize"`
	// Blockhash is the block hash.
	Blockhash string `json:"blockhash"`
	// FeeratePercentiles holds the feerates at the 10th, 25th, 50th, 75th, and 90th percentiles.
	FeeratePercentiles []int `json:"feerate_percentiles"`
	// Height is the height of the block.
	Height int `json:"height"`
	// Ins is the number of inputs (excluding Coinbase).
	Ins int `json:"ins"`
	// MaxFee is the maximum fee in the block.
	MaxFee int `json:"maxfee"`
	// MaxFeeRate is the maximum fee rate in the block (in satoshis per virtual byte).
	MaxFeeRate int `json:"maxfeerate"`
	// MaxTxSize is the maximum transaction size.
	MaxTxSize int `json:"maxtxsize"`
	// MinTxSize is the minimum tx size.
	MinTxSize int `json:"mintxsize"`
	// MedianFee is the median fee in the block.
	MedianFee int `json:"medianfee"`
	// MedianTime is the block mediam time past.
	MedianTime int `json:"mediantime"`
	// MedianTxSize is the median tx size.
	MedianTxSize int `json:"mediantxsize"`
	// Outs is the number of outputs.
	Outs int `json:"outs"`
	// Subsidy is the block subsidy.
	Subsidy int `json:"subsidy"`
	// SegwitTotalSize is the size of all segwit transactions.
	SegwitTotalSize int `json:"swtotal_size"`
	// SegwitTotalWeight is the total weight of all segwit transactions.
	SegwitTotalWeight int `json:"swtotal_weight"`
	// SegwitTxs is the total amount of segwit transactions.
	SegwitTxs int `json:"swtxs"`
	// Time is the block time.
	Time int `json:"time"`
	// TotalOut is the total amount in all outputs (excluding coinbase and thus reward)
	TotalOut int `json:"total_out"`
	// TotalSize is the total size of all non-coinbase transactions.
	TotalSize int `json:"total_size"`
	// TotalWeight is the total weight of all non-coinbase transactions.
	TotalWeight int `json:"total_weight"`
	// TotalFee is the fee total.
	TotalFee int `json:"totalfee"`
	// Txs is the number of transactions (including coinbase).
	Txs int `json:"txs"`
	// UtxoIncrease is the increase/decrease in the nubmer of unspent outputs
	UtxoIncrease int `json:"utxo_increase"`
	// UtxoSizeInc is the increase/decrease in size for the utxo index (not discounting op_return and similiar).
	UtxoSizeInc int `json:"utxo_size_inc"`
}
