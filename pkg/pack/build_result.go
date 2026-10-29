//ff:type feature=pack type=model
//ff:what BuildResult — what a write produced: generation, manifest name, records, shared h/ paths, deletions, objects
package pack

// BuildResult reports a write. Changed is false when nothing differed from
// the previous generation (no objects are written then).
type BuildResult struct {
	Generation uint64
	Manifest   string
	Changed    bool
	Records    int
	Shared     int
	Deleted    int
	Objects    []string
	Bytes      int64
}
