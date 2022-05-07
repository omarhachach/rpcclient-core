package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GenerateBlock mines a block with a set of ordered transactions immediately to a specified address or descriptor.
// txs is either a raw transaction or a txid in the mempool.
func (c *Client) GenerateBlock(output string, txs []string) (*types.GenerateBlockResult, error) {
	var res *types.GenerateBlockResult

	return res, c.SendReq("generateblock", &res, output, txs)
}

// GenerateToAddress mines blocks to a specified address.
// If maxtries is <= 0, will set to default of 1000000.
func (c *Client) GenerateToAddress(nblocks int, address string, maxtries int) ([]string, error) {
	var blocks []string

	if maxtries <= 0 {
		maxtries = 1000000
	}

	return blocks, c.SendReq("generatetoaddress", &blocks, nblocks, address, maxtries)
}

// GenerateToDescriptor is the same as GenerateToAddress, except it uses a descriptor instead of an address.
func (c *Client) GenerateToDescriptor(nblocks int, descriptor string, maxtries int) ([]string, error) {
	var blocks []string

	if maxtries <= 0 {
		maxtries = 1000000
	}

	return blocks, c.SendReq("generatetodescriptor", &blocks, nblocks, descriptor, maxtries)
}
