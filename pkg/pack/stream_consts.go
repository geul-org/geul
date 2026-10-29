package pack

const (
	// DefaultRunEntries is how many H keys a Stream sorts in memory per run (about 120 MB).
	DefaultRunEntries = 1 << 21
	// hEntrySize is the bytes of one H run entry: key 32, object 4, offset 8, length 8, encoding 1.
	hEntrySize = 53
	// streamDir is a Stream's working folder inside the set folder.
	streamDir = ".stream"
)

const (
	// streamStateFile is the checkpoint inside streamDir.
	streamStateFile = "state"
	// streamDataPart and streamIndexPart are the open data pack and index object inside streamDir.
	streamDataPart  = "data.part"
	streamIndexPart = "index.part"
)
