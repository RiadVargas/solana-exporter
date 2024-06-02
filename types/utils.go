package types

type Response struct {
	ID      int    `json:"id"`
	Error   Error  `json:"error,omitempty"`
	Version string `json:"jsonrpc"`
}
