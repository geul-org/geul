//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what TestCheckRange — 206 with the exact Content-Range passes; 200, gzip and a different range are ErrRange
package pack

import (
	"errors"
	"net/http"
	"testing"
)

func TestCheckRange(t *testing.T) {
	resp := func(code int, cr, ce string) *http.Response {
		h := http.Header{}
		h.Set("Content-Range", cr)
		if ce != "" {
			h.Set("Content-Encoding", ce)
		}
		return &http.Response{StatusCode: code, Status: http.StatusText(code), Header: h}
	}
	if err := checkRange(resp(206, "bytes 10-19/100", "identity"), 10, 10); err != nil {
		t.Fatal(err)
	}
	for _, r := range []*http.Response{resp(200, "", ""), resp(206, "bytes 10-19/100", "gzip"), resp(206, "bytes 10-20/100", ""), resp(206, "", "")} {
		if err := checkRange(r, 10, 10); !errors.Is(err, ErrRange) {
			t.Fatal(err)
		}
	}
}
