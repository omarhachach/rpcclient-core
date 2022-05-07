package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GetBlock returns hex-encoded data for the block.
func (c *Client) GetBlock(hash string) (string, error) {
	var hex string

	return hex, c.SendReq("getblock", &hex, hash, 0)
}

// GetBlockVerbose returns an object with information about the block.
func (c *Client) GetBlockVerbose(hash string) (*types.Block, error) {
	var block *types.Block

	return block, c.SendReq("getblock", &block, hash, 1)
}

// GetBlockVerboseTx returns an object with information about the block and the included transactions.
func (c *Client) GetBlockVerboseTx(hash string) (*types.BlockTx, error) {
	var block *types.BlockTx

	return block, c.SendReq("getblock", &block, hash, 2)
}

// GetBlockHash gets the hash of block in best-block-chain at height provided.
func (c *Client) GetBlockHash(height int) (string, error) {
	var blockhash string

	return blockhash, c.SendReq("getblockhash", &blockhash, height)
}

// GetBlockHeader returns a serialized, hex-encoded data for the blockheader.
func (c *Client) GetBlockHeader(blockhash string) (string, error) {
	var blockheader string

	return blockheader, c.SendReq("getblockheader", &blockheader, blockhash, false)
}

// GetBlockHeaderVerbose retrieves a block's header.
func (c *Client) GetBlockHeaderVerbose(blockhash string) (*types.BlockHeader, error) {
	var blockheader *types.BlockHeader

	return blockheader, c.SendReq("getblockheader", &blockheader, blockhash, true)
}

// GetBlockStats computes the per block statstics for a given block.
func (c *Client) GetBlockStats(blockhash string) (*types.BlockStats, error) {
	var blockstats *types.BlockStats

	return blockstats, c.SendReq("getblockstats", &blockstats, blockhash)
}

// GetBlockStatsHeight is the same as GetBlockStats but uses the block height to find the block.
func (c *Client) GetBlockStatsHeight(blockheight int) (*types.BlockStats, error) {
	var blockstats *types.BlockStats

	return blockstats, c.SendReq("getblockstats", &blockstats, blockheight)
}

// PreciousBlock treats a block as if it were received before others with the same work.
func (c *Client) PreciousBlock(blockhash string) error {
	return c.SendReq("preciousblock", new(bool), blockhash)
}
