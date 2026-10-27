//ff:type feature=toulmin type=model
//ff:what ControlCode — the TANGEUL framing controls (section, case boundary, prose)
package toulmin

// ControlCode names a TANGEUL framing control. On the wire each control is its
// own TOULMIN ISSUER-format record (SECTION, CASE-BEGIN, CASE-END, PROSE — GEUL
// t002 TOULMIN); the code values are logical and FROZEN append-only. The old
// stream controls 0x00-0x02 (STREAM_HEADER, SEGMENT_BOUNDARY, CODEBOOK_DECL) are
// retired: the GEUL DOC container and its HEAD Meta replace them.
type ControlCode uint16

// The TANGEUL framing controls. 0x00-0x02 are retired and not reused.
const (
	// ControlSection opens a section; the subcode is a SectionCode (0x03).
	ControlSection ControlCode = 0x03
	// ControlCaseBegin opens a case, carrying the Case TID (0x04).
	ControlCaseBegin ControlCode = 0x04
	// ControlCaseEnd closes a case, carrying the Case TID (0x05).
	ControlCaseEnd ControlCode = 0x05
	// ControlProse marks a prose span (body = SPAN, §7.4) (0x06).
	ControlProse ControlCode = 0x06
)
