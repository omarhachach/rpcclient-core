package types

// BlockChainInfo holds info regarding blockchain processing.
type BlockChainInfo struct {
	Chain                string               `json:"chain"`
	Blocks               int32                `json:"blocks"`
	Headers              int32                `json:"headers"`
	BestBlockHash        string               `json:"bestblockhash"`
	Difficulty           float64              `json:"difficulty"`
	MedianTime           int64                `json:"mediantime"`
	VerificationProgress float64              `json:"verificationprogress"`
	InitialBlockDownload bool                 `json:"initialblockdownload"`
	Chainwork            string               `json:"chainwork"`
	SizeOnDisk           int64                `json:"size_on_disk"`
	Pruned               bool                 `json:"pruned"`
	PruneHeight          int32                `json:"pruneheight"`
	AutomaticPruning     bool                 `json:"automatic_pruning"`
	PruneTargetSize      int                  `json:"prune_target_size"`
	Softforks            map[string]*Softfork `json:"softforks"`
	Warnings             string               `json:"warnings"`
}

// Softfork holds information regarding a soft fork.
type Softfork struct {
	Type   string        `json:"type"`
	Bip9   *SoftforkBip9 `json:"bip9"`
	Height int           `json:"height"`
	Active bool          `json:"active"`
}

// SoftforkBip9 holds status of bip9 softforks (only for "bip9" type).
type SoftforkBip9 struct {
	Status     string             `json:"status"`
	Bit        int                `json:"bit"`
	StartTime  int                `json:"start_time"`
	Timeout    int                `json:"timeout"`
	Since      int                `json:"since"`
	Statistics *SoftforkBip9Stats `json:"statistics"`
}

// SoftforkBip9Stats holds numeric statistics about BIP9 signalling for a softfork (only for "started" status).
type SoftforkBip9Stats struct {
	Period    int  `json:"period"`
	Threshold int  `json:"threshold"`
	Elapsed   int  `json:"elapsed"`
	Count     int  `json:"count"`
	Possible  bool `json:"possible"`
}

// BlockFilter represents the BIP 157 content filter for a particular block.
type BlockFilter struct {
	// Filter is the hex-encoded filter data.
	Filter string `json:"filter"`
	// Header is the hex-encoded filter header.
	Header string `json:"header"`
}

// ChainTip contains information about a chain tip.
type ChainTip struct {
	// Height is the height of the chain tip.
	Height int `json:"height"`
	// Hash is the block hash of the chain tip.
	Hash string `json:"hash"`
	// Branchlen is zero for main chain, otherwise length of branch connection the tip to the main chain.
	Branchlen int `json:"branchlen"`
	// Status is the status of the chain, "active" for the main chain.
	Status string `json:"status"`
}

// ChainTxStats contains statistics about the total number and rate of transactions in the chain.
type ChainTxStats struct {
	Time                   int    `json:"time"`
	Txcount                int    `json:"txcount"`
	WindowFinalBlockHash   string `json:"window_final_block_hash"`
	WindowFinalBlockHeight int    `json:"window_final_block_height"`
	WindowBlockCount       int    `json:"window_block_count"`
	WindowTxCount          int    `json:"window_tx_count"`
	WindowInterval         int    `json:"window_interval"`
	Txrate                 int    `json:"txrate"`
}

// MemoryInfo contains information about memory usage.
type MemoryInfo struct {
	Locked *MemoryInfoLocked `json:"locked"`
}

// MemoryInfoLocked contains information about the locked memory manager.
type MemoryInfoLocked struct {
	Used       int `json:"used"`
	Free       int `json:"free"`
	Total      int `json:"total"`
	Locked     int `json:"locked"`
	ChunksUsed int `json:"chunks_used"`
	ChunksFree int `json:"chunks_free"`
}

// RPCInfo contains detail about the RPC server.
type RPCInfo struct {
	ActiveCommands []*RPCInfoActiveCommand `json:"active_commands"`
	LogPath        string                  `json:"logpath"`
}

// RPCInfoActiveCommand contains info about an active command.
type RPCInfoActiveCommand struct {
	Method   string `json:"method"`
	Duration int    `json:"duration"`
}
