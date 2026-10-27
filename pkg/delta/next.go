//ff:func feature=delta type=helper control=sequence
//ff:what LocalIDs.Next — the NEXT-LOCAL-ID after the IDs handed out so far
package delta

// Next is the next LOCAL ID still free: what the version list records as
// NEXT-LOCAL-ID after this version.
func (l *LocalIDs) Next() uint64 {
	return l.next
}
