//ff:func feature=pack type=writer control=sequence
//ff:what Stream.layerOf — a one-layer list whose root block (kept in the manifest) is root, zstd framed, over n keys
package pack

// layerOf frames root with the codec (checksummed) as the root of the only
// layer, which counts n entries (§4.3 P_LAYERS / H_LAYERS).
func (s *Stream) layerOf(root block, n uint64) ([]layer, error) {
	b, err := s.o.Build.Codec.Compress(encodeBlock(root), nil)
	if err != nil {
		return nil, err
	}
	return []layer{{count: n, enc: encZstd, root: b}}, nil
}
