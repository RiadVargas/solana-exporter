package types

type GetBalanceContext struct {
	Slot int `json:"slot"`
}

type GetBalanceResult struct {
	Context GetBalanceContext `json:"context"`
	Value   uint64            `json:"value"`
}

type GetBalance struct {
	*Response
	Result GetBalanceResult `json:"result"`
}
