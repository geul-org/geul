//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what closeRuns — close the run files of a merge
package pack

// closeRuns closes every run reader's file.
func closeRuns(rs []*runReader) {
	for _, rr := range rs {
		rr.f.Close()
	}
}
