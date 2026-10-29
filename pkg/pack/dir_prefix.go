//ff:func feature=pack type=helper control=sequence
//ff:what dirPrefix — the listing prefix of an fs directory name: "" for ".", "name/" otherwise
package pack

// dirPrefix turns an fs.FS directory name into a path prefix.
func dirPrefix(name string) string {
	if name == "." {
		return ""
	}
	return name + "/"
}
