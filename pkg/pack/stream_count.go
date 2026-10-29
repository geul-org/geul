//ff:func feature=pack type=writer control=sequence
//ff:what Stream.count — tally one stored record in the result: records, compressed ones, h/ shares
package pack

// count adds r to the result counts (BuildResult Records, Compressed, Shared).
func (s *Stream) count(r StreamRecord) {
	s.st.Res.Records++
	if r.Enc == encZstd {
		s.st.Res.Compressed++
	}
	if r.HashOK {
		s.st.Res.Shared++
	}
}
