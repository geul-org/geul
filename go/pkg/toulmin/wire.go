//ff:type feature=toulmin type=model
//ff:what Wire — the per-document field widths a TANGEUL record codec needs (Edge Id, SPAN)
package toulmin

// Wire carries the field widths, in words, that a document declares: ID is the
// Edge Id width (Meta ID-WIDTH) used by every ID/REF field, and Span is the SPAN
// offset/length width declared by the SOURCE record (GEUL t002 TOULMIN 공통
// 필드). Both are 1, 2, 4 or 8.
type Wire struct {
	ID   int
	Span int
}
