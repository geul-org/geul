// Package pack stores a GEUL archive (`.well-known/geul/` paths → bytes) as
// immutable content-named objects: a manifest (`m/<sha256>.gpm`), index
// objects (`o/<sha256>.gpx`) and data packs (`o/<sha256>.gpk`), with
// `latest.txt` the only mutable file (geul-dev files/pack/README.md). The same
// objects serve a local directory (pread) and a CDN (one HTTP Range per read).
package pack

const (
	manifestMagic = "GEULMANI"
	objectMagic   = "GEULPACK"
	formatVersion = 1
	headerSize    = 64

	kindData  = 1 // data pack (.gpk)
	kindIndex = 2 // index object (.gpx)

	secObjects = 0x0001
	secIssuers = 0x0002
	secDicts   = 0x0003
	secNamed   = 0x0004
	secPLayers = 0x0005
	secHLayers = 0x0006

	encRaw  = 0
	encZstd = 1
	encGone = 0xFF // deletion marker in a leaf

	blkPLeaf  = 1
	blkPInner = 2
	blkHLeaf  = 3
	blkHInner = 4

	invP        = 1 // P record
	invH        = 2 // h/-only record
	invNamed    = 3 // named record
	invDict     = 4 // zstd dictionary
	invPDel     = 5 // P deletion
	invNamedDel = 6 // named deletion

	flagHash = 1 // the record also serves `h/<content hash>.geul`

	spaceNamed = 0
	spaceP     = 1
	spaceH     = 2

	latestFile = "latest.txt"
	maxDepth   = 16        // directory depth guard against cycles
	maxBlock   = 128 << 10 // zstd Block_Maximum_Size
	zstdMagic  = 0xFD2FB528
	hashChunk  = 8 << 20 // read size when hashing a whole object

	// DefaultPackSize is the target size of one data pack or index object (1 GiB).
	DefaultPackSize = 1 << 30
	// DefaultPLeaf is the entries per P leaf block.
	DefaultPLeaf = 16384
	// DefaultHLeaf is the entries per H leaf block.
	DefaultHLeaf = 1024
	// DefaultHFanout is the children per H internal block.
	DefaultHFanout = 1024
	// DefaultHRoot is the most children an H root (in the manifest) holds.
	DefaultHRoot = 1024
	// DefaultHLayers is how many H layers stack before they are merged.
	DefaultHLayers = 8
	// DefaultCacheBlocks is how many directory blocks a Reader keeps.
	DefaultCacheBlocks = 4096
)
