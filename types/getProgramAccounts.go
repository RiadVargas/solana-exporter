package types

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetProgramAccountsAccount struct {
	Data       interface{} `json:"data"`
	Executable bool        `json:"executable"`
	Lamports   uint64      `json:"lamports"`
	Owner      string      `json:"owner"`
	Rentepoch  uint64      `json:"rentEpoch"`
}

type GetProgramAccountsResult struct {
	Account GetProgramAccountsAccount `json:"account"`
	Pubkey  string                    `json:"pubkey"`
}

type GetProgramAccounts struct {
	*Response
	Result []GetProgramAccountsResult `json:"result"`
}

type Memcmp struct {
	Offset int    `json:"offset"`
	Bytes  string `json:"bytes"`
}

type Filters struct {
	Memcmp Memcmp `json:"memcmp"`
}

type GetProgramAccountsOpts struct {
	Encoding string    `json:"encoding"`
	Filters  []Filters `json:"filters"`
}
