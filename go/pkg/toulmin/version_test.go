//ff:func feature=toulmin type=model control=sequence
//ff:what TestVersion — the spec version, profile name and codebook version constants
package toulmin

import "testing"

func TestVersion(t *testing.T) {
	if SpecVersion != "geul-t002" || ProfileName != "tangeul" {
		t.Errorf("stamp: got %q %q", SpecVersion, ProfileName)
	}
	if CodebookVersionMajor != 1 || CodebookVersionMinor != 0 {
		t.Errorf("codebook version: got %d.%d, want 1.0", CodebookVersionMajor, CodebookVersionMinor)
	}
}
