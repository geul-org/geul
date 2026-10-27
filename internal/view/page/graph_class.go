//ff:func feature=page type=helper control=sequence
//ff:what graphClass — the graph style class of a TOULMIN record name
package page

// graphClasses groups the TOULMIN record names into graph style classes:
// cases, the three rule roles, definitions, entry points and declarations for
// nodes; attack, supersede, execution and definition relations for edges.
var graphClasses = map[string]string{
	"CASE-DECLARE": "case", "RULE-DECLARE": "rule", "COUNTER-DECLARE": "counter", "EXCEPT-DECLARE": "except",
	"SPEC-DECLARE": "spec", "ENTRY": "entry", "TRIGGER": "entry",
	"NAMESPACE": "decl", "ALIAS": "decl", "FUNC": "decl", "FIELD": "decl", "GROUP": "decl",
	"ATTACKS": "attack", "SUPERSEDES": "supersede",
	"RUN": "run", "DO": "run", "DO-ONCE": "run", "UNDO": "run", "CHECKING": "run",
	"RUNS": "run", "EVALUATES": "run", "TRIGGERS": "run", "UNTIL": "run",
	"REQUIRES": "def", "MEANS": "def", "HAS": "def", "SAME-AS": "def", "IMPORTED-FROM": "def",
}

// graphClass returns the style class of a record name; "other" when unknown.
func graphClass(name string) string {
	if c, ok := graphClasses[name]; ok {
		return c
	}
	return "other"
}
