//ff:func feature=resolve type=resolver control=sequence
//ff:what TestTimeOf — a record's time is the smallest list holding it; past every list is zero
package resolve

import (
	"testing"
	"time"
)

func TestTimeOf(t *testing.T) {
	ts := []listTime{{n: 3, at: time.Unix(30, 0)}, {n: 1, at: time.Unix(10, 0)}, {n: 2, at: time.Unix(20, 0)}}
	if timeOf(ts, 0).Unix() != 10 || timeOf(ts, 1).Unix() != 20 || timeOf(ts, 2).Unix() != 30 {
		t.Fatal("wrong list")
	}
	if !timeOf(ts, 3).IsZero() {
		t.Fatal("past every list: want zero")
	}
}
