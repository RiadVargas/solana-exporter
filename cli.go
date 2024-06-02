package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var (
	MAINNET_RPC             = "https://api.mainnet-beta.solana.com/"
	TESTNET_RPC             = "https://api.testnet.solana.com/"
	DEVNET_RPC              = "https://api.devnet.solana.com/"
	DEFAULT_SCRAPE_INTERVAL = "25"
	STAKE_RENT_RESERVE      = uint64(2282880)

	listenPort     = kingpin.Flag("listenPort", "Solana validator vote account key for narrowing resource-intensive searches").Short('p').Default("8888").String()
	rpcUrl         = kingpin.Flag("rpcUrl", "Solana RPC URL").Short('r').Default(TESTNET_RPC).String()
	scrapeInterval = kingpin.Flag("scrapeInterval", "Scrape interval in seconds").Short('i').Default(DEFAULT_SCRAPE_INTERVAL).Int()
	voteKey        = kingpin.Arg("voteKey", "Solana validator vote account key for narrowing resource-intensive searches").Required().String()
)

func main() {
	kingpin.Parse()

	metrics := Metrics{
		currentEpoch: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "solana_current_epoch",
			Help: "Reports the current epoch",
		}),
		firstSlot: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_first_slot",
			Help: "Reports the first slot height for epoch",
		},
			[]string{"epoch"}),
		lastSlot: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_last_slot",
			Help: "Reports the last expected slot for epoch",
		},
			[]string{"epoch"}),
		currentSlot: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_current_slot_height",
			Help: "Reports the latest slot height for epoch",
		},
			[]string{"epoch"}),
		leaderSlots: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_leadership_slot",
			Help: "Binary metric to foresee leadership slots for identities",
		}, []string{"epoch", "slot", "identity"}),
		//
		identityBalance: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_identity_balance",
			Help: "Latest identity account balance for epoch",
		}, []string{"epoch", "identity", "vote_key"}),
		lastVoteSlot: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_identity_last_vote_slot",
			Help: "Last slot voted by the identity",
		}, []string{"epoch", "identity", "vote_key"}),
		epochCredits: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_vote_epoch_credits",
			Help: "Epoch credits for a given identity and vote key",
		}, []string{"epoch", "identity", "vote_key"}),
		activeStake: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_active_stake",
			Help: "Activate stake for a given identity and vote key",
		}, []string{"epoch", "identity", "vote_key"}),
		pendingStake: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_pending_activation_stake",
			Help: "Pending for activation for a given identity and vote key",
		}, []string{"epoch", "identity", "vote_key"}),
		delinquent: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "solana_validator_delinquent",
			Help: "Binary metric that states if a vote account is delinquent",
		}, []string{"epoch", "identity", "vote_key"}),
	}

	collectorTicker := time.NewTicker(time.Duration(*scrapeInterval) * time.Second)

	go metrics.CollectMetrics(*rpcUrl, *voteKey)
	go func() {
		for {
			<-collectorTicker.C
			metrics.CollectMetrics(*rpcUrl, *voteKey)
		}
	}()

	http.Handle("/*", promhttp.Handler())

	port := fmt.Sprintf(":%s", *listenPort)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to listen on port %s", *listenPort)
	}
}
