//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestPlausibleEmbedded — named Reference with a target, named or ISSUER Value; empty Reference, unnamed codecs and other kinds fail
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestPlausibleEmbedded(t *testing.T) {
	for r, want := range map[*core.Record]bool{
		{Kind: core.KindReference, Codec: core.CodecLocalIdentified, Payload: []core.Word{1}}: true,
		{Kind: core.KindReference, Codec: core.CodecLocalIdentified}:                          false,
		{Kind: core.KindReference, Codec: 0x60, Payload: []core.Word{1}}:                      false,
		{Kind: core.KindValue, Codec: core.CodecUTF8TextEven}:                                 true,
		{Kind: core.KindValue, Codec: core.CodecIssuer1Word, Issuer: []core.Word{1}}:          true,
		{Kind: core.KindValue, Codec: 0x30}:                                                   false,
		{Kind: core.KindEdge, Codec: 0x00}:                                                    false,
	} {
		if plausibleEmbedded(*r) != want {
			t.Fatalf("%+v: want %v", *r, want)
		}
	}
}
