package util

// EthereumStatDump is a Struct used to fetch JSON blockcypher's REST api
type EthereumStatDump struct {
	Height         int64  `json:"height"`
	Hash           string `json:"hash"`
	Time           string `json:"time"`
	PreviousHash   string `json:"previous_hash"`
	PeerCount      int64  `json:"peer_count"`
	HighGasPrice   int64  `json:"high_gas_price"`
	LowGasPrice    int64  `json:"low_gas_price"`
	LastForkHeight int64  `json:"last_fork_height"`
	LastForkHash   string `json:"last_fork_hash"`
}

func (dump *EthereumStatDump) fetch() {

}
