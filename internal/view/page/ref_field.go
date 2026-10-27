//ff:type feature=page type=model
//ff:what refField — one field of a Reference payload: its name and word count
package page

// refField names N words of a Reference payload (GEUL t002 OVERVIEW Reference).
type refField struct {
	Name string
	N    int
}
