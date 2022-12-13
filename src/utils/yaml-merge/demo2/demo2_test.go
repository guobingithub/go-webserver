package demo2

import "testing"

func Test_readAndMergeConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readAndMergeConfig()
		})
	}
}
