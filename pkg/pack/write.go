//ff:func feature=pack type=writer control=sequence
//ff:what Write — add one generation to the pack set in out: new records in new packs, layers, manifest, then latest.txt
package pack

import "context"

// Write records b as the next generation of the pack set in directory out
// (`latest.txt`, `m/`, `o/`; created when missing). Objects are only ever
// added; latest.txt moves last. When nothing changes no generation is made.
func Write(ctx context.Context, out string, b Batch, o BuildOptions) (BuildResult, error) {
	return writeTo(ctx, out, b, o, osFS{})
}
