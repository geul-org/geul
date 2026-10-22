//ff:func feature=geulmd type=helper control=selection
//ff:what toulminName — the GEULMD keyword of a TOULMIN direct codec or GEUL format number
package geulmd

import "github.com/geul-org/geul/go/pkg/toulmin"

// toulminNames maps every TOULMIN record type to its spec name (GEUL t002
// TOULMIN 코덱). Direct codecs are keyed by codec, formats by format number.
var toulminNames = map[toulmin.CodecRef]string{
	{Codec: toulmin.CodecCaseDeclare}: "CASE-DECLARE", {Codec: toulmin.CodecPredicateDeclare}: "PREDICATE-DECLARE",
	{Codec: toulmin.CodecSpecDeclare}: "SPEC-DECLARE", {Codec: toulmin.CodecRuleDeclare}: "RULE-DECLARE",
	{Codec: toulmin.CodecCounterDeclare}: "COUNTER-DECLARE", {Codec: toulmin.CodecExceptDeclare}: "EXCEPT-DECLARE",
	{Codec: toulmin.CodecAttacks}: "ATTACKS", {Codec: toulmin.CodecSupersedes}: "SUPERSEDES",
	{Codec: toulmin.CodecRun}: "RUN", {Codec: toulmin.CodecDo}: "DO", {Codec: toulmin.CodecDoOnce}: "DO-ONCE",
	{Codec: toulmin.CodecUndo}: "UNDO", {Codec: toulmin.CodecChecking}: "CHECKING",
	{Codec: toulmin.CodecEntry}: "ENTRY", {Codec: toulmin.CodecTrigger}: "TRIGGER",
}

// toulminFormats maps the GEUL ISSUER format numbers to their spec names.
var toulminFormats = map[uint16]string{
	toulmin.FormatSource: "SOURCE", toulmin.FormatText: "TEXT", toulmin.FormatSection: "SECTION",
	toulmin.FormatCaseBegin: "CASE-BEGIN", toulmin.FormatCaseEnd: "CASE-END", toulmin.FormatProse: "PROSE",
	toulmin.FormatNamespace: "NAMESPACE", toulmin.FormatFunc: "FUNC", toulmin.FormatField: "FIELD",
	toulmin.FormatAlias: "ALIAS", toulmin.FormatGroup: "GROUP", toulmin.FormatRequires: "REQUIRES",
	toulmin.FormatMeans: "MEANS", toulmin.FormatHas: "HAS", toulmin.FormatRuns: "RUNS",
	toulmin.FormatEvaluates: "EVALUATES", toulmin.FormatTriggers: "TRIGGERS", toulmin.FormatUntil: "UNTIL",
	toulmin.FormatSameAs: "SAME-AS", toulmin.FormatImportedFrom: "IMPORTED-FROM", toulmin.FormatLabel: "LABEL",
}

// toulminName returns the keyword of ref, or "" when ref is no TOULMIN record type.
func toulminName(ref toulmin.CodecRef) string {
	switch ref.Format {
	case 0:
		return toulminNames[ref]
	default:
		return toulminFormats[ref.Format]
	}
}
