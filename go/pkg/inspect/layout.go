//ff:func feature=inspect type=decoder control=sequence
//ff:what Layout — a record's words split into named head and payload parts
package inspect

// Layout splits e's words into parts in order: the head (TAG WORD, Length
// extension, Issuer and format number) and then the payload, field by field
// when e's typed content gives the layout, otherwise as generically as the
// spec allows. Concatenating every part's words gives e.Words back.
func Layout(e Entry) []Part {
	return append(headParts(e), payloadParts(e)...)
}
