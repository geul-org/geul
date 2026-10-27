//ff:func feature=resolve type=resolver control=sequence
//ff:what HTTP.Fetch — GET a path; 404 and 410 are ErrNotFound, other non-200 statuses are errors
package resolve

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Fetch GETs path under Base. 404 and 410 mean the source does not have it.
func (h HTTP) Fetch(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strings.TrimSuffix(h.Base, "/")+"/"+wellKnown+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	c := h.Client
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
		return nil, ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geul/resolve: %s: %s", req.URL, resp.Status)
	}
	return io.ReadAll(resp.Body)
}
