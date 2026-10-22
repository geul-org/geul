//ff:type feature=core type=model
//ff:what Document — a parsed GEUL DOC: its HEAD Meta records and BODY records
package core

// Document is one parsed GEUL document (GEUL t002 OVERVIEW 구조): the HEAD Meta
// records in order and the BODY records in order. The DOC-*, BODY and END
// structure records are validated and consumed by SplitDocument, not kept.
type Document struct {
	Head []Record
	Body []Record
}
