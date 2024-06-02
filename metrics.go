package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	currentEpoch    prometheus.Gauge
	currentSlot     *prometheus.GaugeVec
	firstSlot       *prometheus.GaugeVec
	lastSlot        *prometheus.GaugeVec
	leaderSlots     *prometheus.GaugeVec
	identityBalance *prometheus.GaugeVec
	lastVoteSlot    *prometheus.GaugeVec
	epochCredits    *prometheus.GaugeVec
	activeStake     *prometheus.GaugeVec
	pendingStake    *prometheus.GaugeVec
	delinquent      *prometheus.GaugeVec
}
