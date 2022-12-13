package demo2

import "testing"

func TestMergeConfigIntoDefault(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "xxx01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeConfigIntoDefault()
		})
	}
}
