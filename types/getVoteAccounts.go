package types

const (
	VALIDATOR_DELINQUENT = 1
	VALIDATOR_ACTIVE     = 0
)

type VoteAccount struct {
	Commission       int     `json:"commission"`
	EpochVoteAccount bool    `json:"epochVoteAccount"`
	EpochCredits     [][]int `json:"epochCredits"`
	NodePubkey       string  `json:"nodePubkey"`
	LastVote         int     `json:"lastVote"`
	ActivatedStake   int     `json:"activatedStake"`
	VotePubkey       string  `json:"votePubkey"`
}

type GetVoteAccountsResult struct {
	Current    []VoteAccount `json:"current"`
	Delinquent []VoteAccount `json:"delinquent"`
}

type GetVoteAccounts struct {
	*Response
	Result GetVoteAccountsResult `json:"result"`
}

type GetVoteAccountsOpts struct {
	VotePubKey string `json:"votePubkey"`
}
