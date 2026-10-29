//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterWorkers — Zstd's workers, 1 for Stored
package pack

import "testing"

func TestWriterWorkers(t *testing.T) {
	if (&writer{o: BuildOptions{Codec: &Zstd{Workers: 5}}}).workers() != 5 || (&writer{o: BuildOptions{Codec: Stored{}}}).workers() != 1 {
		t.Fatal("workers")
	}
}
