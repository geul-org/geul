//ff:func feature=pack type=helper control=sequence
//ff:what TestSortedKeys — ascending
package pack

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	if k := sortedKeys(keySet{"b": {}, "a": {}}); !reflect.DeepEqual(k, []string{"a", "b"}) {
		t.Fatal(k)
	}
}
