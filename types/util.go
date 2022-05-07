package types

// EstimateSmartFeeResult is the result of the estimatesmartfee call.
type EstimateSmartFeeResult struct {
	FeeRate float64  `json:"feerate"`
	Errors  []string `json:"errors"`
	Blocks  int      `json:"blocks"`
}
