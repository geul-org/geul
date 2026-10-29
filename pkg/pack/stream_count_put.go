//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Count — how many P records the stream holds (after a resume: as of the checkpoint)
package pack

// Count returns the number of P records put; after ResumeStream it is the
// count at the checkpoint, so the caller knows how many to skip.
func (s *Stream) Count() int64 {
	return s.st.Count
}
