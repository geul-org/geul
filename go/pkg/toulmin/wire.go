//ff:type feature=toulmin type=model
//ff:what Wire — the per-document field width a TANGEUL record codec needs (SPAN)
package toulmin

// Wire carries the field width, in words, that a document declares: Span is
// the SPAN offset/length width declared by the SOURCE record (GEUL t002 TOULMIN
// 공통 필드), 1, 2, 4 or 8. Edge Ids and REFs need no width: they are embedded
// LOCAL Reference records sized by their own Length.
type Wire struct {
	Span int
}
