//ff:func feature=codecspec type=helper control=iteration dimension=1
//ff:what Spec.Merge — s with o's codecs added where s declares none
package codecspec

// Merge returns a Spec holding every codec of s and the codecs of o that s
// does not declare: s wins.
func (s Spec) Merge(o Spec) Spec {
	out := Spec{}
	for k, v := range o {
		out[k] = v
	}
	for k, v := range s {
		out[k] = v
	}
	return out
}
