//ff:func feature=resolve type=decoder control=sequence
//ff:what TestParseVersionList — round trip with blank lines; a bad line and out-of-order versions fail
package resolve

import (
	"reflect"
	"testing"
)

func TestParseVersionList(t *testing.T) {
	list := []VersionEntry{{Version: 1}, {Version: 3, Deprecated: true}}
	got, err := parseVersionList([]byte("\n" + formatVersionList(list) + "\n"))
	if err != nil || !reflect.DeepEqual(got, list) {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := parseVersionList([]byte("junk\n")); err == nil {
		t.Fatal("bad line: want error")
	}
	if _, err := parseVersionList([]byte(formatVersionList([]VersionEntry{{Version: 2}, {Version: 2}}))); err == nil {
		t.Fatal("order: want error")
	}
}
