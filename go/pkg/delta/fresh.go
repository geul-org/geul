//ff:func feature=delta type=encoder control=sequence
//ff:what LocalIDs.fresh — the next new LOCAL ID from the auto increment
package delta

// fresh hands out the next new LOCAL ID and advances the auto increment.
func (l *LocalIDs) fresh() uint64 {
	id := l.next
	l.next++
	l.used[id] = true
	return id
}
