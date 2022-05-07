package types

// MempoolTransaction holds information regarding a transaction in the mempool.
type MempoolTransaction struct {
	Vsize             int                     `json:"vsize"`
	Weight            int                     `json:"weight"`
	Time              int                     `json:"time"`
	Height            int                     `json:"height"`
	DescendantCount   int                     `json:"descendantcount"`
	DescendantSize    int                     `json:"descendantsize"`
	AncestorCount     int                     `json:"ancestorcount"`
	AncestorSize      int                     `json:"ancestorsize"`
	WTxId             string                  `json:"wtxid"`
	Fees              *MempoolTransactionFees `json:"fees"`
	Depends           []string                `json:"depends"`
	SpentBy           []string                `json:"spentby"`
	Bip125Replaceable bool                    `json:"bip125-replaceable"`
	Unbroadcast       bool                    `json:"unbroadcast"`
}

// MempoolTransactionFees holds information about a MempoolTransaction's fees.
type MempoolTransactionFees struct {
	Base       float64 `json:"base"`
	Modified   float64 `json:"modified"`
	Ancestor   float64 `json:"ancestor"`
	Descendant float64 `json:"descendant"`
}

// MempoolInfo contains info about the active state of the transaction mempool.
type MempoolInfo struct {
	Loaded           bool    `json:"loaded"`
	Size             int     `json:"size"`
	Bytes            int     `json:"bytes"`
	Usage            int     `json:"usage"`
	MaxMempool       int     `json:"maxmempool"`
	MempoolMinFee    float64 `json:"mempoolminfee"`
	MinRelayTxFee    float64 `json:"minrelaytxfee"`
	UnbroadcastCount int     `json:"unbroadcastcount"`
}
