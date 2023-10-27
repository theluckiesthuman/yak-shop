package implementation

import (
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/require"
)

func TestCalculateStock(t *testing.T) {
	f, err := os.Open("../../../data/herd.xml")
	require.NoError(t, err)
	defer f.Close()
	fq, err := NewFileQuery(f)
	require.NoError(t, err)
	tests := []struct {
		name string
		T    int
	}{
		{
			name: "T=13",
			T:    13,
		},
		{
			name: "T=14",
			T:    14,
		},
		{
			name: "T=1000",
			T:    1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stock, err := fq.CalculateStock(tt.T)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, stock)
		})
	}
}

func TestCalculateAge(t *testing.T) {
	f, err := os.Open("../../../data/herd.xml")
	require.NoError(t, err)
	defer f.Close()

	fq, err := NewFileQuery(f)
	require.NoError(t, err)
	tests := []struct {
		name string
		T    int
	}{
		{
			name: "T=13",
			T:    13,
		},
		{
			name: "T=14",
			T:    14,
		},
		{
			name: "T=1000",
			T:    1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			herd, err := fq.CalculateAge(13)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, herd)
		})
	}
}
