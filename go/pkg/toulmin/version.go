//ff:type feature=toulmin type=model
//ff:what version stamp — GEUL spec version, TANGEUL profile name and codebook version constants
package toulmin

// The version stamp a TANGEUL document declares in its HEAD (GEUL t002 TOULMIN
// 문서): Meta SPEC-VERSION carries SpecVersion and Meta PROFILE carries
// "<ProfileName>-<major>.<minor>". A reader rejects a different SpecVersion or
// an unknown codebook major; minor drift is tolerated (additive changes only).
//
// The codebook major moved 0→1 when the binary moved from the Free-area
// `.geul` framing to the GEUL t002 DOC container (`.geul`): every record's
// bit layout changed, so old readers must not accept the new files.
const (
	// SpecVersion is the GEUL specification this build writes and reads.
	SpecVersion = "geul-t002"
	// ProfileName names the TANGEUL profile in Meta PROFILE.
	ProfileName = "tangeul"

	// CodebookVersionMajor is this codebook's major version.
	CodebookVersionMajor uint8 = 1
	// CodebookVersionMinor is this codebook's minor version.
	CodebookVersionMinor uint8 = 0
)
