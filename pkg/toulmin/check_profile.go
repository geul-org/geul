//ff:func feature=toulmin type=decoder control=sequence
//ff:what CheckProfile — accept a Meta PROFILE text of this profile and codebook major, any minor
package toulmin

import "fmt"

// CheckProfile validates a Meta PROFILE text: it must be
// "<ProfileName>-<major>.<minor>" with this build's codebook major. A different
// profile name or major is rejected (the record layouts may differ); any minor
// is accepted because minor changes are additive.
func CheckProfile(s string) error {
	var name string
	var major, minor uint8
	if _, err := fmt.Sscanf(s, "%7s-%d.%d", &name, &major, &minor); err != nil || name != ProfileName || s != fmt.Sprintf("%s-%d.%d", name, major, minor) {
		return fmt.Errorf("geul/toulmin: unknown profile %q", s)
	}
	if major != CodebookVersionMajor {
		return fmt.Errorf("geul/toulmin: profile %q has codebook major %d (this build is %d)", s, major, CodebookVersionMajor)
	}
	return nil
}
