package main

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/riadvargas/solana-exporter/types"
	"github.com/rs/zerolog/log"
)

func (m Metrics) RegisterPendingStake(rpcUrl string, currentEpoch string, voteAccount types.VoteAccount) {
	labels := prometheus.Labels{"epoch": currentEpoch, "identity": voteAccount.NodePubkey, "vote_key": voteAccount.VotePubkey}

	programAccounts := GetStakeProgramAccounts(rpcUrl, voteAccount.VotePubkey)

	var delegatedStake uint64
	for _, stake := range programAccounts.Result {
		delegatedStake += stake.Account.Lamports - STAKE_RENT_RESERVE
	}

	pendingStake := delegatedStake - uint64(voteAccount.ActivatedStake)
	m.pendingStake.With(labels).Set(float64(pendingStake))
}

func (m Metrics) RegisterLeaderSchedule(rpcUrl string, currentEpoch string, identity string, epochInfo types.GetEpochInfo) {
	leaderSlots := GetLeaderSchedule(rpcUrl, identity)

	m.currentEpoch.Set(float64(epochInfo.Result.Epoch))

	initialSlot := float64(epochInfo.Result.AbsoluteSlot - epochInfo.Result.SlotIndex)
	lastSlot := initialSlot + float64(epochInfo.Result.SlotsInEpoch) - 1

	m.firstSlot.With(prometheus.Labels{"epoch": currentEpoch}).Set(initialSlot)
	m.lastSlot.With(prometheus.Labels{"epoch": currentEpoch}).Set(lastSlot)

	for k, leaders := range leaderSlots.Result {
		for _, s := range leaders {
			m.leaderSlots.With(prometheus.Labels{"epoch": currentEpoch, "slot": fmt.Sprintf("%0.f", s+initialSlot), "identity": k}).Set(1)
		}
	}
}

func (m Metrics) RegisterVoteAccount(rpcUrl string, currentEpoch string, voteAccount types.VoteAccount, delinquent int) {
	labels := prometheus.Labels{"epoch": currentEpoch, "identity": voteAccount.NodePubkey, "vote_key": voteAccount.VotePubkey}

	m.lastVoteSlot.With(labels).Set(float64(voteAccount.LastVote))
	m.activeStake.With(labels).Set(float64(voteAccount.ActivatedStake))
	m.delinquent.With(labels).Set(float64(delinquent))

	for _, epochCredit := range voteAccount.EpochCredits {
		epoch := epochCredit[0]
		currentCredits := epochCredit[1]
		previousCredits := epochCredit[2]
		gainedCredits := (currentCredits - previousCredits)

		go m.epochCredits.With(prometheus.Labels{"epoch": strconv.Itoa(epoch), "identity": voteAccount.NodePubkey, "vote_key": voteAccount.VotePubkey}).Set(float64(gainedCredits))
	}

	balance := GetBalance(rpcUrl, voteAccount.NodePubkey)

	m.identityBalance.With(labels).Set(float64(balance.Result.Value))
}

func (m Metrics) RegisterVoteAccounts(currentEpoch string, voteAccounts []types.VoteAccount) {
	labels := prometheus.Labels{"epoch": currentEpoch}

	var clusterTotalCredits int
	var eligibleValidators int
	for _, v := range voteAccounts {
		if len(v.EpochCredits) > 0 {
			epochCredit := v.EpochCredits[len(v.EpochCredits)-1]
			creditEpoch := strconv.Itoa(epochCredit[0])

			if creditEpoch == currentEpoch {
				currentCredits := epochCredit[1]
				previousCredits := epochCredit[2]
				gainedCredits := (currentCredits - previousCredits)

				clusterTotalCredits += gainedCredits
				eligibleValidators++
			}
		}
	}

	averageCredits := (clusterTotalCredits / eligibleValidators)

	m.averageCredits.With(labels).Set(float64(averageCredits))
}

func (m Metrics) CollectMetrics(rpcUrl string, voteKey string) {
	log.Info().Msg("starting a new collection job")

	epochInfo := GetEpochInfo(rpcUrl)
	currentEpoch := fmt.Sprintf("%d", epochInfo.Result.Epoch)

	voteAccountResponse := GetVoteAccount(rpcUrl, voteKey)

	for _, voteAccount := range append(voteAccountResponse.Result.Current, voteAccountResponse.Result.Delinquent...) {
		go m.RegisterVoteAccount(rpcUrl, currentEpoch, voteAccount, types.VALIDATOR_ACTIVE)

		// @TODO: run just once per epoch
		go m.RegisterLeaderSchedule(rpcUrl, currentEpoch, voteAccount.NodePubkey, epochInfo)
		go m.RegisterPendingStake(rpcUrl, currentEpoch, voteAccount)
		// --- @TODO: run just once per epoch
	}

	voteAccountsResponse := GetVoteAccounts(rpcUrl)

	go m.RegisterVoteAccounts(currentEpoch, append(voteAccountsResponse.Result.Current, voteAccountsResponse.Result.Delinquent...))

	m.currentSlot.With(prometheus.Labels{"epoch": currentEpoch}).Set(float64(epochInfo.Result.AbsoluteSlot))
}
