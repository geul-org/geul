//ff:type feature=pack type=model
//ff:what BuildResult — what a write produced: generation, manifest name, records, shared h/ paths, deletions, objects, compressed records, new dictionary
package pack

// BuildResult reports a write. Changed is false when nothing differed from
// the previous generation (no objects are written then). Compressed counts the
// records stored as zstd; Dict is the ID of a dictionary trained by this
// write (0: none).
type BuildResult struct {
	Generation uint64
	Manifest   string
	Changed    bool
	Records    int
	Shared     int
	Deleted    int
	Objects    []string
	Bytes      int64
	Compressed int
	Dict       uint32
}
