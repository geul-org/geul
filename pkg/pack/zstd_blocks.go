//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what zstdBlocks — decode a frame's blocks up to the last one; returns content and the bytes after
package pack

// zstdBlocks decodes blocks until Last_Block. Blocks are at most
// min(window, 128 KiB).
func zstdBlocks(b []byte, window uint64) ([]byte, []byte, error) {
	limit := min(window, maxBlock)
	out := []byte{}
	for {
		var last bool
		var err error
		out, b, last, err = zstdBlock(out, b, limit)
		if err != nil {
			return nil, nil, err
		}
		if last {
			return out, b, nil
		}
	}
}
