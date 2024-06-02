package types

type GetEpochInfoConfig struct {
	Commitment     string `json:"commitment"`
	MinContextSlot int    `json:"minContextSlot"`
}

type GetEpochInfoResult struct {
	AbsoluteSlot     int `json:"absoluteSlot"`
	BlockHeight      int `json:"blockHeight"`
	Epoch            int `json:"epoch"`
	SlotIndex        int `json:"slotIndex"`
	SlotsInEpoch     int `json:"slotsInEpoch"`
	TransactionCount int `json:"transactionCount"`
}

type GetEpochInfo struct {
	*Response
	Result GetEpochInfoResult `json:"result"`
}
