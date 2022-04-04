package rpcclient

import (
	"github.com/omarhachach/rpcclient-core/types"
)

// GetMempoolAncestors gets a list of transaction ID's for the in-mempool ancestors of the provided txid.
func (c *Client) GetMempoolAncestors(txid string) ([]string, error) {
	var txids []string

	return txids, c.SendReq("getmempoolancestors", &txids, txid, false)
}

// GetMempoolAncestorsVerbose is like GetMempoolAncestors but will map the transaction ID's to detail objects.
func (c *Client) GetMempoolAncestorsVerbose(txid string) (map[string]*types.MempoolTransaction, error) {
	var txs map[string]*types.MempoolTransaction

	return txs, c.SendReq("getmempoolancestors", &txs, txid, true)
}

// GetMempoolDescendants gets a list of transaction ID's for the in-mempool descendants of the provided txid.
func (c *Client) GetMempoolDescendants(txid string) ([]string, error) {
	var txids []string

	return txids, c.SendReq("getmempooldescendants", &txids, txid, false)
}

// GetMempoolDescendantsVerbose is like GetMempoolDescendants but will map the transaction ID's to detail objects.
func (c *Client) GetMempoolDescendantsVerbose(txid string) (map[string]*types.MempoolTransaction, error) {
	var txs map[string]*types.MempoolTransaction

	return txs, c.SendReq("getmempooldescendants", &txs, txid, true)
}

// GetMempoolEntry retrieves the mempool data for a given transaction. (Txid must be in mempool).
func (c *Client) GetMempoolEntry(txid string) (*types.MempoolTransaction, error) {
	var tx *types.MempoolTransaction

	return tx, c.SendReq("getmempoolentry", &tx, txid)
}

// GetMempoolInfo returns details on the active state of the transaction memory pool.
func (c *Client) GetMempoolInfo() (*types.MempoolInfo, error) {
	var info *types.MempoolInfo

	return info, c.SendReq("getmempoolinfo", &info)
}

// GetRawMempool returns a list of txids in the mempool.
func (c *Client) GetRawMempool() ([]string, error) {
	var txids []string

	return txids, c.SendReq("getrawmempool", &txids, false, false)
}

// GetRawMempoolVerbose is like GetRawMempool but will map the transaction ID's to detail objects.
func (c *Client) GetRawMempoolVerbose() (map[string]*types.MempoolTransaction, error) {
	var txs map[string]*types.MempoolTransaction

	return txs, c.SendReq("getrawmempool", &txs, true, false)
}

// SaveMempool dumps the mempool to disk.
func (c *Client) SaveMempool() error {
	return c.SendReq("savemempool", new(bool))
}
