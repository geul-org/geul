//ff:type feature=toulmin type=model
//ff:what EdgePacket — a decoded LINK record: Edge Id, relation, src, dst, flags, optional gate and SPAN
package toulmin

// EdgePacket is a decoded LINK record (GEUL t002 TOULMIN LINK). The record codec
// gives Prop; DO-ONCE additionally sets F_ONCE. The dst form is decided by Prop
// via PropDstKind: do/undo carry a FuncRef (DstFunc), every other
// relation a REF (DstTID). Gate is present iff F_GATE.
type EdgePacket struct {
	ID      TID
	Prop    PropCode
	Src     TID
	DstTID  TID
	DstFunc *FuncRef
	Flags   EdgeFlags
	Gate    *Gate
	Span    Span
}
