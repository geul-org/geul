//ff:func feature=pack type=resolver control=sequence
//ff:what HTTPStore.get — one GET of an object with extra headers; 404 and 410 are ErrNotFound
package pack

import (
	"context"
	"net/http"
	"strings"
)

// get sends a GET for name with the given header pairs.
func (s HTTPStore) get(ctx context.Context, name string, hdr ...string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimSuffix(s.Base, "/")+"/"+name, nil)
	if err != nil {
		return nil, err
	}
	for i := 0; i+1 < len(hdr); i += 2 {
		req.Header.Set(hdr[i], hdr[i+1])
	}
	c := s.Client
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
		resp.Body.Close()
		return nil, notFound(name)
	}
	return resp, nil
}
