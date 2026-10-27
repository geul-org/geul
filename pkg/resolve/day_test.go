//ff:func feature=resolve type=helper control=sequence
//ff:what day — test fixture: midnight UTC of a day count from 1970-01-01
package resolve

import (
	"time"
)

func day(n int64) time.Time {
	return time.Unix(n*86400, 0).UTC()
}
