package types

type GetLeaderScheduleOpts struct {
	Commitment string `json:"commitment"`
	Identity   string `json:"identity"`
}

type GetLeaderScheduleResult struct {
	*Response
	Result map[string][]float64 `json:"result"`
}
