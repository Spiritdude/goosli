package test

import (
	. "github.com/l1va/goosli/primitives"
	"github.com/l1va/goosli/slicers"
	"github.com/l1va/goosli/slicers/vip"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestVip_Slice(t *testing.T) {
	cases := []struct {
		in     string
		layers int
	}{
		{
			in:     "bokal",
			layers: 800,
		},
		{
			in:     "leyka",
			layers: 500,
		},
		{
			in:     "odn",
			layers: 1060,
		},
		{
			in:     "odn2",
			layers: 954,
		},
		{
			in:     "odnbez",
			layers: 722,
		},
		{
			in:     "odnzak",
			layers: 850,
		},
		{
			in:     "ruchka",
			layers: 811,
		},
		{
			in:     "shtutser",
			layers: 250,
		},
		{
			in:     "truba",
			layers: 757,
		},
	}

	for _, row := range cases {
		t.Run(row.in, func(t *testing.T) {
			mesh, err := LoadSTL("../test_models/" + row.in + ".stl")
			if err != nil {
				log.Fatal("failed to load mesh: ", err)
			}
			sett := slicers.Settings{
				LayerHeight:   0.2,
				WallThickness: 1.2,
				FillDensity:   20,
				Nozzle:        0.4,
			}

			gcd := vip.Slice(mesh, sett)

			require.Equal(t, row.layers, gcd.LayersCount)
		})
	}
}