//ff:func feature=notation type=encoder control=sequence
//ff:what TestNamedEdge — declared codecs by fields, undeclared token by token; a bad field and an unknown name fail
package notation

import "testing"

func TestNamedEdge(t *testing.T) {
	if e, err := namedEdge(Edge{Label: "2"}, []string{"TRIPLE", "P1", "Q1", "Q2", "normal"}, Context{Spec: tripleSpec(false)}); err != nil || e.Label != "2" || len(e.Rest) != 7 {
		t.Fatalf("%+v %v", e, err)
	}
	if e, err := namedEdge(Edge{}, []string{"BIND", "#1", `"a"`}, Context{}); err != nil || len(e.Rest) != 4 {
		t.Fatalf("%+v %v", e, err)
	}
	if _, err := namedEdge(Edge{}, []string{"BIND", "word"}, Context{}); err == nil {
		t.Fatal("bad field: want error")
	}
	if _, err := namedEdge(Edge{}, []string{"NOPE"}, Context{}); err == nil {
		t.Fatal("unknown: want error")
	}
}
