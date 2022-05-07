package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// EstimateSmartFee estimates the approximate fee per kilobyte needed for a transaction to begin with confTarget blocks.
func (c *Client) EstimateSmartFee(confTarget int, estimateMode *types.EstimateMode) (*types.EstimateSmartFeeResult, error) {
	var res *types.EstimateSmartFeeResult

	if estimateMode != nil {
		return res, c.SendReq("estimatesmartfee", &res, confTarget, *estimateMode)
	}

	return res, c.SendReq("estimatesmartfee", &res, confTarget)
}
