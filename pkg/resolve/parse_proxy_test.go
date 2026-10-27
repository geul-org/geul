//ff:func feature=resolve type=decoder control=sequence
//ff:what TestParseProxy — `|` marks fall-through on the item before it; a bad item fails the list
package resolve

import (
	"testing"
)

func TestParseProxy(t *testing.T) {
	es, err := parseProxy("file:///a | https://b.example,off")
	if err != nil || len(es) != 3 || !es[0].Fallthrough || es[1].Fallthrough || !es[2].Off {
		t.Fatalf("got %+v %v", es, err)
	}
	if _, err := parseProxy("off,,registry"); err == nil {
		t.Fatal("empty item: want error")
	}
}
