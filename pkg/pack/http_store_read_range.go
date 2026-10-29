//ff:func feature=pack type=resolver control=sequence
//ff:what HTTPStore.ReadRange — one `Range: bytes=a-b` request; 206, Content-Range and exact length required
package pack

import (
	"context"
	"fmt"
	"io"
)

// ReadRange fetches bytes [off, off+n) of name in one request.
func (s HTTPStore) ReadRange(ctx context.Context, name string, off, n uint64) ([]byte, error) {
	if n == 0 {
		return []byte{}, nil
	}
	resp, err := s.get(ctx, name, "Range", fmt.Sprintf("bytes=%d-%d", off, off+n-1), "Accept-Encoding", "identity")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := checkRange(resp, off, n); err != nil {
		return nil, err
	}
	b, err := io.ReadAll(io.LimitReader(resp.Body, int64(n)+1))
	if err == nil && uint64(len(b)) != n {
		err = fmt.Errorf("%w: %s: body %d bytes, want %d", ErrRange, name, len(b), n)
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}
