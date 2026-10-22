//ff:type feature=toulmin type=model
//ff:what TextPacket — a TEXT (or SOURCE) record: Edge Id plus verbatim UTF-8 bytes
package toulmin

// TextPacket is a TEXT record — an Edge Id and a UTF-8 Value with the verbatim
// bytes (GEUL t002 TOULMIN ISSUER 형식). Labels, paths and values ride here and
// are referenced by the Edge Id. The SOURCE record (the authored document)
// shares this type. Bytes are stored verbatim (no Unicode normalization).
type TextPacket struct {
	TID   TID
	Bytes []byte
}
