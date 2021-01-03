package util

import (
	"encoding/json"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

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

// Fetch fetches and deserializes Json from blockcypher's REST api i
func (dump *EthereumStatDump) Fetch() error {
	r, err := client.Get("https://api.blockcypher.com/v1/eth/main")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(dump)
}
