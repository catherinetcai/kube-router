package bgp

import (
	"testing"

	"github.com/cloudnativelabs/kube-router/v2/pkg/options"
	gobgpapi "github.com/osrg/gobgp/v4/api"
	"github.com/stretchr/testify/assert"
)

func TestBuildPeerBfd(t *testing.T) {
	tests := []struct {
		name     string
		peer     BFDConfig
		expected *gobgpapi.BfdPeerConfig
	}{
		{
			name: "bfd not enabled returns nil",
			peer: BFDConfig{Enabled: new(bool), Port: new(uint32(5000))},
		},
		{
			name: "fields not set are set with defaults",
			peer: BFDConfig{
				Enabled:              new(true),
				DetectionMultiplier:  new(uint32(5)),
				DesiredMinTxInterval: new(uint32(2000000)),
			},
			expected: &gobgpapi.BfdPeerConfig{
				Enabled:                  true,
				Port:                     3784,
				DetectionMultiplier:      5,
				DesiredMinimumTxInterval: 2000000,
				RequiredMinimumReceive:   options.DefaultBFDRequiredMinRxInterval,
			},
		},
		{
			name: "all fields set, not overridden by defaults",
			peer: BFDConfig{
				Enabled:               new(true),
				Port:                  new(uint32(3785)),
				DetectionMultiplier:   new(uint32(2)),
				DesiredMinTxInterval:  new(uint32(2000000)),
				RequiredMinRxInterval: new(uint32(2000000)),
			},
			expected: &gobgpapi.BfdPeerConfig{
				Enabled:                  true,
				Port:                     3785,
				DetectionMultiplier:      2,
				DesiredMinimumTxInterval: 2000000,
				RequiredMinimumReceive:   2000000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BuildPeerBfd(tt.peer)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
