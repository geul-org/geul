//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Meta — the caller's progress saved with the last checkpoint (nil before any)
package pack

// Meta returns what the caller passed to the last Checkpoint; after
// ResumeStream it tells the caller where it stopped.
func (s *Stream) Meta() []byte {
	return s.st.Meta
}
