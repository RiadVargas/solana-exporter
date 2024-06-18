package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/riadvargas/solana-exporter/types"
)

type Request struct {
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Version string        `json:"jsonrpc"`
}

func request(rpcUrl string, payload *Request) ([]byte, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	log.Debug().RawJSON("request", b).Msg("starting an RPC request")

	client := &http.Client{}
	buffer := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", rpcUrl, buffer)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	return body, nil
}

func GetBalance(rpcUrl string, identity string) types.GetBalance {
	payload := &Request{
		ID:     1,
		Method: "getBalance",
		Params: []interface{}{
			identity,
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetBalance
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetVoteAccount(rpcUrl string, voteKey string) types.GetVoteAccounts {
	payload := &Request{
		ID:     1,
		Method: "getVoteAccounts",
		Params: []interface{}{
			&types.GetVoteAccountsOpts{
				VotePubKey: voteKey,
			},
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetVoteAccounts
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetVoteAccounts(rpcUrl string) types.GetVoteAccounts {
	payload := &Request{
		ID:      1,
		Method:  "getVoteAccounts",
		Params:  []interface{}{},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetVoteAccounts
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetEpochInfo(rpcUrl string) types.GetEpochInfo {
	payload := &Request{
		ID:     1,
		Method: "getEpochInfo",
		Params: []interface{}{
			types.GetEpochInfoConfig{
				Commitment: "finalized",
			},
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetEpochInfo
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetLeaderSchedule(rpcUrl string, identity string) types.GetLeaderScheduleResult {
	payload := &Request{
		ID:     1,
		Method: "getLeaderSchedule",
		Params: []interface{}{
			nil,
			types.GetLeaderScheduleOpts{
				Identity: identity,
			},
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetLeaderScheduleResult
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetStakeProgramAccounts(rpcUrl string, voteKey string) types.GetProgramAccounts {
	payload := &Request{
		ID:     1,
		Method: "getProgramAccounts",
		Params: []interface{}{
			"Stake11111111111111111111111111111111111111",
			&types.GetProgramAccountsOpts{
				Encoding: "base64",
				Filters: []types.Filters{
					{
						Memcmp: types.Memcmp{
							Offset: 124,
							Bytes:  voteKey,
						},
					},
				},
			},
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetProgramAccounts
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}

func GetAccountInfo(rpcUrl string, identity string) types.GetAccountInfo {
	payload := &Request{
		ID:     1,
		Method: "getAccountInfo",
		Params: []interface{}{
			identity,
			&types.GetProgramAccountsOpts{
				Encoding: "jsonParsed",
			},
		},
		Version: "2.0",
	}

	body, err := request(rpcUrl, payload)
	if err != nil {
		panic(err)
	}

	var resp types.GetAccountInfo
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Fatal().Err(err)
	}

	return resp
}
