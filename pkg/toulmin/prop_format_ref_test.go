//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestPropFormatRef — every format-number relation maps to its format; a direct-code relation does not
package toulmin

import "testing"

func TestPropFormatRef(t *testing.T) {
	want := map[PropCode]uint16{
		PropRequires: FormatRequires, PropMeans: FormatMeans, PropHas: FormatHas,
		PropRuns: FormatRuns, PropEvaluates: FormatEvaluates, PropTriggers: FormatTriggers,
		PropUntil: FormatUntil, PropSameAs: FormatSameAs, PropImportedFrom: FormatImportedFrom,
		PropLabel: FormatLabel,
	}
	for p, f := range want {
		if ref, err := propFormatRef(p); err != nil || ref.Codec != 0x7C || ref.Format != f {
			t.Fatalf("0x%04X: %+v %v", uint16(p), ref, err)
		}
	}
	if _, err := propFormatRef(PropAttacks); err == nil {
		t.Fatal("attacks has a direct code, not a format")
	}
}
