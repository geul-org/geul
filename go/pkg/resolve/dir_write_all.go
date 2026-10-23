//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.writeAll — store several paths, version list documents last, stopping at the first failure
package resolve

import "sort"

// writeAll stores every path of files in name order, version list documents
// last, so a reader never sees a list naming a version whose file is missing.
func (d Dir) writeAll(files map[string][]byte) error {
	paths := make([]string, 0, len(files))
	for p := range files {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	sort.SliceStable(paths, func(i, j int) bool { return !isListPath(paths[i]) && isListPath(paths[j]) })
	for _, p := range paths {
		if err := d.write(p, files[p]); err != nil {
			return err
		}
	}
	return nil
}
