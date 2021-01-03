package util

// EthereumStatDump is a Struct used to fetch JSON blockcypher's REST api
type EthereumStatDump struct {
	Height         int64
	Hash           string
	Time           string
	PreviousHash   string
	PeerCount      int64
	HighGasPrice   int64
	LowGasPrice    int64
	LastForkHeight int64
	LastForkHash   string
}

func (dump *EthereumStatDump) fetch() {
	
}
