package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GetBestBlockHash returns the hash of the best (tip) block in the most-work fully-validated chain.
func (c *Client) GetBestBlockHash() (string, error) {
	var hash string

	return hash, c.SendReq("getbestblockhash", &hash)
}

// GetBlockChainInfo returns various state info regarding blockchain processing.
func (c *Client) GetBlockChainInfo() (*types.BlockChainInfo, error) {
	var info *types.BlockChainInfo

	return info, c.SendReq("getblockchaininfo", &info)
}

// GetBlockCount returns the number of blocks in the longest block chain.
func (c *Client) GetBlockCount() (int64, error) {
	var count int64

	return count, c.SendReq("getblockcount", &count)
}

// GetBlockFilter retrieves the BIP 157 content filter for a particular block.
func (c *Client) GetBlockFilter(blockhash, filtertype string) (*types.BlockFilter, error) {
	var filter *types.BlockFilter

	return filter, c.SendReq("getblockfilter", &filter, blockhash, filtertype)
}

// GetChainTips returns information about all known tips in the block tree, including the main chain as well as
// orphaned branches.
func (c *Client) GetChainTips() ([]*types.ChainTip, error) {
	var tips []*types.ChainTip

	return tips, c.SendReq("getchaintips", &tips)
}

// GetChainTxStats computes statistics about the total number and rate of transactions in the chain.
func (c *Client) GetChainTxStats(nblocks int, blockhash string) (*types.ChainTxStats, error) {
	var txStats *types.ChainTxStats

	return txStats, c.SendReq("getchaintxstats", &txStats, nblocks, blockhash)
}

// GetDifficulty returns the proof-of-work difficulty as a multiple of the minimum difficulty.
func (c *Client) GetDifficulty() (float64, error) {
	var diff float64

	return diff, c.SendReq("getdifficulty", &diff)
}

// PruneBlockchain prunes the blockchain up to the given height.
func (c *Client) PruneBlockchain(height int) (int, error) {
	var lastPruned int

	return lastPruned, c.SendReq("pruneblockchain", &lastPruned, height)
}

// VerifyChain verifies the blockchain database.
func (c *Client) VerifyChain(level int) (bool, error) {
	var verified bool

	return verified, c.SendReq("verifychain", &verified, level)
}

// GetMemoryInfo returns an object containing information about memory usage.
func (c *Client) GetMemoryInfo() (*types.MemoryInfo, error) {
	var info *types.MemoryInfo

	return info, c.SendReq("getmemoryinfo", &info, "stats")
}

// GetMemoryInfoMalloc returns an XML string describing low-level heap state. (Only available if node is compiled with
// glibc 2.10+).
func (c *Client) GetMemoryInfoMalloc() (string, error) {
	var malloc string

	return malloc, c.SendReq("getmemoryinfo", &malloc, "mallocinfo")
}

// GetRPCInfo returns details about the RPC server.
func (c *Client) GetRPCInfo() (*types.RPCInfo, error) {
	var info *types.RPCInfo

	return info, c.SendReq("getrpcinfo", &info)
}

// Uptime returns the total uptime of the server in seconds.
func (c *Client) Uptime() (int, error) {
	var uptime int

	return uptime, c.SendReq("uptime", &uptime)
}

// Stop requests a graceful shutdown of the node.
func (c *Client) Stop() error {
	return c.SendReq("stop", "")
}
