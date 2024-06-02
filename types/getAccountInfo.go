package types

type GetAccountInfoOpts struct {
	Encoding string `json:"encoding"`
}

type GetAccountInfoContext struct {
	Slot int `json:"slot"`
}

type AuthorizedVoters struct {
	AuthorizedVoter string `json:"authorizedVoter"`
	Epoch           int    `json:"epoch"`
}

type EpochCredits struct {
	Credits         string `json:"credits"`
	Epoch           int    `json:"epoch"`
	PreviousCredits string `json:"previousCredits"`
}

type LastTimestamp struct {
	Slot      int `json:"slot"`
	Timestamp int `json:"timestamp"`
}

type Votes struct {
	ConfirmationCount uint64 `json:"confirmationCount"`
	Slot              uint64 `json:"slot"`
}

type GetAccountInfoDataParsedInfo struct {
	AuthorizedVoters     []AuthorizedVoters `json:"authorizedVoters"`
	AuthorizedWithdrawer string             `json:"authorizedWithdrawer"`
	Commission           uint64             `json:"commission"`
	EpochCredits         []EpochCredits     `json:"epochCredits"`
	LastTimestamp        LastTimestamp      `json:"lastTimestamp"`
	NodePubkey           string             `json:"nodePubkey"`
	PriorVoters          []interface{}      `json:"priorVoters"`
	RootSlot             uint64             `json:"rootSlot"`
	Votes                []Votes            `json:"votes"`
}

type GetAccountInfoDataParsed struct {
	Info GetAccountInfoDataParsedInfo `json:"info"`
}

type GetAccountInfoData struct {
	Parsed GetAccountInfoDataParsed `json:"parsed"`
}

type GetAccountInfoValue struct {
	Data       GetAccountInfoData `json:"data"`
	Executable bool               `json:"executable"`
	Lamports   uint64             `json:"lamports"`
	Owner      string             `json:"owner"`
	RentEpoch  uint64             `json:"rentEpoch"`
	Space      uint64             `json:"space"`
}

type GetAccountInfoResult struct {
	Context GetAccountInfoContext `json:"context"`
	Value   GetAccountInfoValue   `json:"value"`
}

type GetAccountInfo struct {
	*Response
	Result GetAccountInfoResult `json:"result"`
}
