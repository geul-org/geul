//ff:func feature=core type=helper control=selection
//ff:what publicationTail — the words after the version ID for each version list codec
package core

// publicationTail returns how many words follow the version ID: 16 (content
// hash) for PUBLICATION, 0 for DEPRECATION, 1 (reason) for TOMBSTONE. The bool
// is false for any other codec.
func publicationTail(codec uint8) (int, bool) {
	switch codec {
	case EdgePublication:
		return 16, true
	case EdgeDeprecation:
		return 0, true
	case EdgeTombstone:
		return 1, true
	default:
		return 0, false
	}
}
