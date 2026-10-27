//ff:type feature=build type=model
//ff:what Count — how much a build wrote: documents and directories
package build

// Count is the size of a build: Docs .geul documents (each copied and given a
// viewer page) and Dirs directories (each given a listing).
type Count struct {
	Docs int
	Dirs int
}
