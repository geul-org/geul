//ff:func feature=toulmin type=encoder control=sequence
//ff:what ProfileString — the Meta PROFILE text "tangeul-<major>.<minor>" of this build
package toulmin

import "fmt"

// ProfileString returns the Meta PROFILE text this build writes:
// "<ProfileName>-<CodebookVersionMajor>.<CodebookVersionMinor>".
func ProfileString() string {
	return fmt.Sprintf("%s-%d.%d", ProfileName, CodebookVersionMajor, CodebookVersionMinor)
}
