//ff:func feature=core type=decoder control=selection
//ff:what placeDocRecord — route one record into HEAD or BODY, handling BODY/END and Meta placement rules
package core

import "errors"

// placeDocRecord applies one record to st (GEUL t002 OVERVIEW 구조): BODY may
// appear once and ends the HEAD; END closes the document; Meta records belong to
// the HEAD only and Reference/Value/Edge records to the BODY only. A non-Meta
// record before any HEAD record starts an implicit BODY (BODY omitted because
// the HEAD is empty). Any record after END is an error.
func placeDocRecord(st *docState, r Record) error {
	if st.ended {
		return errors.New("geul/core: record after END")
	}
	isMeta := r.Kind == KindMeta
	switch {
	case isMeta && r.Codec == MetaEnd:
		st.ended = true
	case isMeta && r.Codec == MetaBody:
		if st.inBody {
			return errors.New("geul/core: BODY appears twice or after body records")
		}
		st.inBody = true
	case isMeta && r.Codec <= MetaDoc4Word:
		return errors.New("geul/core: nested DOC record (documents do not nest)")
	case isMeta:
		if st.inBody {
			return errors.New("geul/core: Meta record in BODY")
		}
		st.doc.Head = append(st.doc.Head, r)
	case !st.inBody && len(st.doc.Head) > 0:
		return errors.New("geul/core: non-Meta record in HEAD (missing BODY)")
	default:
		st.inBody = true
		st.doc.Body = append(st.doc.Body, r)
	}
	return nil
}
