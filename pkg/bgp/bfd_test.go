package bgp

import (
	"testing"

	gobgpapi "github.com/osrg/gobgp/v4/api"
)

func TestBuildPeerBfd(t *testing.T) {
	tests := []struct {
		name    string
		peer    BFDConfig
		enable  bool
		wantNil bool
		want    *gobgpapi.BfdPeerConfig
	}{
		{
			name:    "cluster-wide disabled returns nil",
			peer:    BFDConfig{Enabled: true, Port: 5000},
			enable:  false,
			wantNil: true,
		},
		{
			name:    "per-peer not enabled returns nil",
			peer:    BFDConfig{Enabled: false, Port: 5000},
			enable:  true,
			wantNil: true,
		},
		{
			name:   "per-peer enabled with no override uses cluster defaults",
			peer:   BFDConfig{Enabled: true},
			enable: true,
			want: &gobgpapi.BfdPeerConfig{
				Enabled:                  true,
				Port:                     3784,
				DetectionMultiplier:      3,
				DesiredMinimumTxInterval: 1000000,
				RequiredMinimumReceive:   1000000,
			},
		},
		{
			name: "per-peer partial override merges with defaults",
			peer: BFDConfig{
				Enabled:              true,
				DetectionMultiplier:  5,
				DesiredMinTxInterval: 2000000,
			},
			enable: true,
			want: &gobgpapi.BfdPeerConfig{
				Enabled:                  true,
				Port:                     3784,
				DetectionMultiplier:      5,
				DesiredMinimumTxInterval: 2000000,
				RequiredMinimumReceive:   1000000,
			},
		},
		{
			name: "per-peer full override replaces all defaults",
			peer: BFDConfig{
				Enabled:               true,
				Port:                  3785,
				DetectionMultiplier:   2,
				DesiredMinTxInterval:  2000000,
				RequiredMinRxInterval: 2000000,
			},
			enable: true,
			want: &gobgpapi.BfdPeerConfig{
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
			got := BuildPeerBfd(tt.peer, tt.enable)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("expected nil, got %+v", got)
				}
				return
			}
			if got == nil {
				t.Fatal("expected non-nil BfdPeerConfig")
			}
			if got.Port != tt.want.Port ||
				got.DetectionMultiplier != tt.want.DetectionMultiplier ||
				got.DesiredMinimumTxInterval != tt.want.DesiredMinimumTxInterval ||
				got.RequiredMinimumReceive != tt.want.RequiredMinimumReceive {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
