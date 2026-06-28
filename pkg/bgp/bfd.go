package bgp

import (
	gobgpapi "github.com/osrg/gobgp/v4/api"
)

// BFDConfig holds BFD timer configuration. As a per-peer override on
// PeerConfig, Enabled gates whether the override applies; when false, the
// cluster-wide defaults are used. When true, any zero-valued timer field
// is filled from the cluster defaults.
type BFDConfig struct {
	Enabled               bool   `yaml:"enabled"`
	DesiredMinTxInterval  uint32 `yaml:"desired_min_tx_interval"`
	RequiredMinRxInterval uint32 `yaml:"required_min_rx_interval"`
	DetectionMultiplier   uint32 `yaml:"detection_multiplier"`
	Port                  uint32 `yaml:"port"`
}

// BuildPeerBfd returns the gobgpapi.BfdPeerConfig to set on a peer, or
// nil if BFD should not be enabled for this peer. peerCfg.Enabled is the
// per-peer gate; enableBFD is the cluster-wide gate. Both must be true.
// Zero-valued timer fields in peerCfg are filled from clusterDefaults.
func BuildPeerBfd(peerCfg BFDConfig, enableBFD bool, clusterDefaults BFDConfig) *gobgpapi.BfdPeerConfig {
	if !enableBFD || !peerCfg.Enabled {
		return nil
	}

	port := peerCfg.Port
	if port == 0 {
		port = clusterDefaults.Port
	}
	multiplier := peerCfg.DetectionMultiplier
	if multiplier == 0 {
		multiplier = clusterDefaults.DetectionMultiplier
	}
	tx := peerCfg.DesiredMinTxInterval
	if tx == 0 {
		tx = clusterDefaults.DesiredMinTxInterval
	}
	rx := peerCfg.RequiredMinRxInterval
	if rx == 0 {
		rx = clusterDefaults.RequiredMinRxInterval
	}

	return &gobgpapi.BfdPeerConfig{
		Enabled:                  true,
		Port:                     port,
		DetectionMultiplier:      multiplier,
		DesiredMinimumTxInterval: tx,
		RequiredMinimumReceive:   rx,
	}
}
