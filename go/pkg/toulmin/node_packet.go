//ff:type feature=toulmin type=model
//ff:what NodePacket — a decoded NODE record: Edge Id, type, header fields, flag-gated slots and SPAN
package toulmin

// NodePacket is a decoded NODE record (GEUL t002 TOULMIN NODE). The record codec
// gives Type (and, for the three rule codecs, Role); the header word packs Role
// (bits 15-14), Level (bits 10-9), the Claim object kind (bits 8-7) and Flags
// (bits 6-0). Each optional field is present exactly when its flag is set:
// Func↔F_FUNC, Qual↔F_QUAL, Check↔F_CHECK, Claim↔F_CLAIM, Verdict↔F_VERDICT,
// With↔F_WITH, Grounds↔F_GROUNDS; F_STRIKE has no field. With holds TEXT Edge
// Ids of the `with` term names and Grounds those of the `grounded on` facts. Verdict/Qual are VFix values written as SIGNED-DECIMAL Values.
type NodePacket struct {
	TID     TID
	Type    NodeType
	Role    RoleCode
	Level   LevelCode
	Flags   NodeFlags
	Func    *FuncRef
	Qual    *VFix
	Check   *TID
	Claim   *Claim
	Verdict *VFix
	With    []TID
	Grounds []TID
	Span    Span
}
