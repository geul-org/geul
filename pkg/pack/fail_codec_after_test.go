//ff:func feature=pack type=encoder control=sequence
//ff:what failCodecAfter — a failAfter codec allowing n calls
package pack

// failCodecAfter returns a codec that succeeds n times, then fails.
func failCodecAfter(n int) Codec {
	return failAfter{left: &n}
}
