//ff:type feature=toulmin type=model
//ff:what Claim — the proposition a node carries: subject, predicate, object, mode
package toulmin

// Claim is the proposition a node carries — its cache key. It encodes as the
// subject REF, the predicate code word, the object (REF / UNSIGNED-INTEGER Value
// / SPAN, selected by Kind) and the mode code word (GEUL t002 TOULMIN NODE).
// Kind is not written inside the block but in the NODE header word. The
// type-ladder fallback is a descriptive claim with Subject=Predicate=0,
// Kind=ObjSpan and ObjSpan anchoring the label.
type Claim struct {
	Subject   TID
	Predicate uint16
	Kind      ObjKind
	ObjTID    TID
	ObjNum    uint32
	ObjSpan   Span
	Mode      uint16
}
