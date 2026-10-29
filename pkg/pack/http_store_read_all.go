//ff:func feature=pack type=resolver control=sequence
//ff:what HTTPStore.ReadAll — GET a whole object; anything but 200 is an error
package pack

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ReadAll fetches name.
func (s HTTPStore) ReadAll(ctx context.Context, name string) ([]byte, error) {
	resp, err := s.get(ctx, name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geul/pack: GET %s: %s", name, resp.Status)
	}
	return io.ReadAll(resp.Body)
}
