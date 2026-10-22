//ff:type feature=toulmin type=model
//ff:what ControlPacket — a decoded framing record: SECTION, CASE-BEGIN, CASE-END or PROSE
package toulmin

// ControlPacket is a decoded framing record (GEUL t002 TOULMIN ISSUER 형식):
// SECTION (Arg = section code), CASE-BEGIN / CASE-END (Arg = the case's Edge
// Id), and PROSE (Span = the prose source span). Span is non-nil exactly for
// PROSE. The document-level controls of the old stream (STREAM_HEADER,
// SEGMENT_BOUNDARY) are gone: the DOC container and its HEAD replace them.
type ControlPacket struct {
	ID   TID
	Code ControlCode
	Arg  uint64
	Span *Span
}
