//ff:func feature=pack type=encoder control=sequence
//ff:what mustUnzstd — decode a zstd frame without dictionary or fail the test
package pack

import "testing"

// mustUnzstd decodes frame with the plain decoder.
func mustUnzstd(t *testing.T, frame []byte) []byte {
	t.Helper()
	b, err := plainDecoder().DecodeAll(frame, nil)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
