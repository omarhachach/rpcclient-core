package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GetBlockTemplate returns data needed to construct a block to work on.
// If template is nil, will use default.
func (c *Client) GetBlockTemplate(template *types.BlockTemplateRequest) (*types.BlockTemplate, error) {
	var tmplt *types.BlockTemplate

	if template != nil {
		return tmplt, c.SendReq("getblocktemplate", &tmplt, template)
	}

	return tmplt, c.SendReq("getblocktemplate", &tmplt)
}

// GetMiningInfo returns mining-related information.
func (c *Client) GetMiningInfo() (*types.MiningInfo, error) {
	var info *types.MiningInfo

	return info, c.SendReq("getmininginfo", &info)
}

// GetNetworkHashPS returns the estimated network hashes per second based on the last nblocks.
func (c *Client) GetNetworkHashPS(nblocks, height int) (int, error) {
	var hsps int

	return hsps, c.SendReq("getnetworkhashps", nblocks, height)
}

// PrioritiseTransaction accepts the transaction into mined blocks at a higher (or lower) priority.
func (c *Client) PrioritiseTransaction(txid string, feeDelta int) (bool, error) {
	var res bool

	return res, c.SendReq("prioritisetransaction", &res, txid, feeDelta)
}

// SubmitBlock submits a new block to the network.
func (c *Client) SubmitBlock(hexdata string) error {
	return c.SendReq("submitblock", new(string), hexdata)
}

// SubmitHeader decodes the hexdata as a header and submits it as a candidate chain tip if valid.
func (c *Client) SubmitHeader(hexdata string) error {
	return c.SendReq("submitblock", new(string), hexdata)
}
