package bgp

import (
	"github.com/cloudnativelabs/kube-router/v2/pkg/options"
	gobgpapi "github.com/osrg/gobgp/v4/api"
)

type BFDConfig struct {
	Enabled               bool   `yaml:"enabled"`
	DesiredMinTxInterval  uint32 `yaml:"desired_min_tx_interval"`
	RequiredMinRxInterval uint32 `yaml:"required_min_rx_interval"`
	DetectionMultiplier   uint32 `yaml:"detection_multiplier"`
	Port                  uint32 `yaml:"port"`
}

func BuildPeerBfd(peerCfg BFDConfig, enableBFD bool) *gobgpapi.BfdPeerConfig {
	if !enableBFD || !peerCfg.Enabled {
		return nil
	}

	port := peerCfg.Port
	if port == 0 {
		port = options.DefaultBFDPort
	}
	multiplier := peerCfg.DetectionMultiplier
	if multiplier == 0 {
		multiplier = options.DefaultBFDDetectionMultiplier
	}
	tx := peerCfg.DesiredMinTxInterval
	if tx == 0 {
		tx = options.DefaultBFDDesiredMinTxInterval
	}
	rx := peerCfg.RequiredMinRxInterval
	if rx == 0 {
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
