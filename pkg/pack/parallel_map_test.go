//ff:func feature=pack type=helper control=sequence
//ff:what TestParallelMap — results in index order whatever the workers; the first error by index; nothing to do
package pack

import (
	"errors"
	"testing"
)

func TestParallelMap(t *testing.T) {
	out, err := parallelMap(100, 7, func(i int) (int, error) { return i * i, nil })
	if err != nil || len(out) != 100 || out[99] != 99*99 || out[3] != 9 {
		t.Fatal(err)
	}
	_, err = parallelMap(10, 3, func(i int) (int, error) {
		if i >= 4 {
			return 0, errors.New(string(rune('a' + i)))
		}
		return i, nil
	})
	if err == nil || err.Error() != "e" {
		t.Fatal(err)
	}
	if out, err := parallelMap(0, 4, func(int) (int, error) { return 0, nil }); err != nil || len(out) != 0 {
		t.Fatal(out, err)
	}
}
