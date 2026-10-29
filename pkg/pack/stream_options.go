//ff:type feature=pack type=writer
//ff:what StreamOptions — a Stream's block sizes, codec, dictionary samples, Issuers and run size
package pack

// StreamOptions tune a Stream. Build gives the pack and block sizes, the
// codec (nil is Zstd at its defaults) and the clock, as for Write; its
// Verify, Check and Fresh are not used. Issuers lists the Issuers (canonical
// hex) of the P records, in ISSUERS order. Samples are records a dictionary
// is trained on (even positions) and tried on (odd positions) before the
// first record, as Write samples a batch; fewer than 32 train none.
// RunEntries is how many H keys are held before a sorted run is spilled to
// disk (0 is DefaultRunEntries). Created is the manifest's created_at (Unix
// seconds; 0 takes Build.Now when the stream starts).
type StreamOptions struct {
	Build      BuildOptions
	Issuers    []string
	Samples    [][]byte
	RunEntries int
	Created    int64
}
