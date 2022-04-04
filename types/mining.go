package types

// BlockTemplateRequest is
type BlockTemplateRequest struct {
	Mode         string   `json:"mode,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	Rules        []string `json:"rules"`
}

// BlockTemplate contains info about a mining block template.
type BlockTemplate struct {
	Version                  int                         `json:"version"`
	Rules                    []string                    `json:"rules"`
	VbAvailable              map[string]int              `json:"vbavailable"`
	VbRequried               int                         `json:"vbrequried"`
	PreviousBlockhash        string                      `json:"previousblockhash"`
	Transactions             []*BlockTemplateTransaction `json:"transactions"`
	CoinbaseAux              map[string]string           `json:"coinbaseaux"`
	CoinbaseValue            int                         `json:"coinbasevalue"`
	LongpollId               string                      `json:"longpollid"`
	Target                   string                      `json:"target"`
	MinTime                  int                         `json:"mintime"`
	Mutable                  []string                    `json:"mutable"`
	NonceRange               string                      `json:"noncerange"`
	SigopLimit               int                         `json:"sigoplimit"`
	SizeLimit                int                         `json:"sizelimit"`
	WeightLimit              int                         `json:"weightlimit"`
	CurTime                  int                         `json:"curtime"`
	Bits                     string                      `json:"bits"`
	Height                   int                         `json:"height"`
	DefaultWitnessCommitment string                      `json:"default_witness_commitment"`
}

// BlockTemplateTransaction contens of a non-coinbase transaction. Used in BlockTemplate.
type BlockTemplateTransaction struct {
	Data    string `json:"data"`
	Txid    string `json:"txid"`
	Hash    string `json:"hash"`
	Depends []int  `json:"depends"`
	Fee     int    `json:"fee"`
	SigOps  int    `json:"sigops"`
	Weight  int    `json:"weight"`
}

// MiningInfo contains mining-related information.
type MiningInfo struct {
	Blocks             int    `json:"blocks"`
	CurrentBlockWeight int    `json:"currentblockweight"`
	CurrentBlockTx     int    `json:"currentblocktx"`
	Difficulty         int    `json:"difficulty"`
	NetworkHashPS      int    `json:"networkhashps"`
	PooledTx           int    `json:"pooledtx"`
	Chain              string `json:"chain"`
	Warnings           string `json:"warnings"`
}
