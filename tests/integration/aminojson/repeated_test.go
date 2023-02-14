package aminojson

import (
	gogoproto "github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/proto"

	"cosmossdk.io/x/tx/aminojson"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	gogopb "github.com/cosmos/cosmos-sdk/tests/integration/aminojson/internal/gogo/testpb"
	pulsarpb "github.com/cosmos/cosmos-sdk/tests/integration/aminojson/internal/pulsar/testpb"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepeatedFields(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	aj := aminojson.NewAminoJSON()

	cases := map[string]struct {
		gogo   gogoproto.Message
		pulsar proto.Message
		fails  bool
	}{
		"unsupported_empty_sets": {
			gogo:   &gogopb.TestRepeatedFields{},
			pulsar: &pulsarpb.TestRepeatedFields{},
			fails:  true,
		},
		"unsupported_empty_sets_are_set": {
			gogo: &gogopb.TestRepeatedFields{
				NullableDontOmitempty: []*gogopb.Streng{{Value: "foo"}},
				NonNullableOmitempty:  []gogopb.Streng{{Value: "foo"}},
			},
			pulsar: &pulsarpb.TestRepeatedFields{
				NullableDontOmitempty: []*pulsarpb.Streng{{Value: "foo"}},
				NonNullableOmitempty:  []*pulsarpb.Streng{{Value: "foo"}},
			},
		},
		"unsupported_nullable": {
			gogo:   &gogopb.TestNullableFields{},
			pulsar: &pulsarpb.TestNullableFields{},
			fails:  true,
		},
		"unsupported_nullable_set": {
			gogo: &gogopb.TestNullableFields{
				NullableDontOmitempty: &gogopb.Streng{Value: "foo"},
				NonNullableOmitempty:  gogopb.Streng{Value: "foo"},
			},
			pulsar: &pulsarpb.TestNullableFields{
				NullableDontOmitempty: &pulsarpb.Streng{Value: "foo"},
				NonNullableOmitempty:  &pulsarpb.Streng{Value: "foo"},
			},
		},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			gogoBz, err := cdc.MarshalJSON(tc.gogo)
			require.NoError(t, err)
			pulsarBz, err := aj.MarshalAmino(tc.pulsar)
			require.NoError(t, err)

			fmt.Printf("  gogo: %s\npulsar: %s\n", string(gogoBz), string(pulsarBz))

			if tc.fails {
				require.NotEqual(t, string(gogoBz), string(pulsarBz))
			} else {
				require.Equal(t, string(gogoBz), string(pulsarBz))
			}
		})
	}
}
