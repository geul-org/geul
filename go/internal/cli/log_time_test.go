//ff:func feature=cli type=helper control=sequence
//ff:what TestLogTime — UTC time; zero is -
package cli

import (
	"testing"
	"time"
)

func TestLogTime(t *testing.T) {
	if logTime(time.Unix(0, 0)) != "1970-01-01 00:00:00" || logTime(time.Time{}) != "-" {
		t.Fatal("logTime")
	}
}
