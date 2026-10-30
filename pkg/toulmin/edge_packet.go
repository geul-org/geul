//ff:type feature=toulmin type=model
//ff:what EdgePacket — a decoded LINK record: Edge Id, relation, src, dst, flags, optional gate and force, SPAN
package toulmin

// EdgePacket is a decoded LINK record (GEUL t002 TOULMIN LINK). The record codec
// gives Prop; DO-ONCE additionally sets F_ONCE. The dst form is decided by Prop
// via PropDstKind: do/undo carry a FuncRef (DstFunc), every other
// relation a REF (DstTID). Gate is present iff F_GATE; Force (the attack
// force in [0, 1], ATTACKS only) iff F_FORCE. Without F_FORCE the force is 1.
type EdgePacket struct {
	ID      TID
	Prop    PropCode
	Src     TID
	DstTID  TID
	DstFunc *FuncRef
	Flags   EdgeFlags
	Gate    *Gate
	Force   *VFix
	Span    Span
}
