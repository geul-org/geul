//ff:func feature=cli type=helper control=sequence
//ff:what harness.must — test fixture: run a command that has to succeed; its output
package cli

// must runs args and fails the test on an error.
func (h *harness) must(args ...string) string {
	h.t.Helper()
	out, err := h.run("", args...)
	if err != nil {
		h.t.Fatalf("geul %v: %v", args, err)
	}
	return out
}
