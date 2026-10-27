//ff:func feature=toulmin type=encoder control=sequence
//ff:what doCodec — DO or DO-ONCE by the once flag
package toulmin

// doCodec returns DO-ONCE for a `do … once` edge and DO otherwise.
func doCodec(once bool) CodecRef {
	if once {
		return CodecRef{Codec: CodecDoOnce}
	}
	return CodecRef{Codec: CodecDo}
}
