package bgp

import (
	"github.com/cloudnativelabs/kube-router/v2/pkg/options"
	gobgpapi "github.com/osrg/gobgp/v4/api"
)

type BFDConfig struct {
	Enabled               *bool   `yaml:"enabled,omitempty"`
	Port                  *uint32 `yaml:"port"`
	DesiredMinTxInterval  *uint32 `yaml:"desired_min_tx_interval"`
	DetectionMultiplier   *uint32 `yaml:"detection_multiplier"`
	RequiredMinRxInterval *uint32 `yaml:"required_min_rx_interval"`
}

func BuildPeerBfd(peerCfg BFDConfig) *gobgpapi.BfdPeerConfig {
	if peerCfg.Enabled != nil && !(*peerCfg.Enabled) {
		return nil
	}

	var port uint32
	if peerCfg.Port != nil {
		port = *peerCfg.Port
	} else {
		port = options.DefaultBFDPort
	}
	var multiplier uint32
	if peerCfg.DetectionMultiplier != nil {
		multiplier = *peerCfg.DetectionMultiplier
	} else {
		multiplier = options.DefaultBFDDetectionMultiplier
	}
	var tx uint32
	if peerCfg.DesiredMinTxInterval != nil {
		tx = *peerCfg.DesiredMinTxInterval
	} else {
		tx = options.DefaultBFDDesiredMinTxInterval
	}
	var rx uint32
	if peerCfg.RequiredMinRxInterval != nil {
		rx = *peerCfg.RequiredMinRxInterval
	} else {
		rx = options.DefaultBFDRequiredMinRxInterval
	}

	return &gobgpapi.BfdPeerConfig{
		Enabled:                  true,
		Port:                     port,
		DetectionMultiplier:      multiplier,
		DesiredMinimumTxInterval: tx,
		RequiredMinimumReceive:   rx,
	}
}
