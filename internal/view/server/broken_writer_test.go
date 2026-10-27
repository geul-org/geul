//ff:type feature=server type=helper
//ff:what brokenWriter — test ResponseWriter that accepts headers but fails body writes
package server

import "net/http/httptest"

// brokenWriter accepts headers but fails every body write.
type brokenWriter struct{ *httptest.ResponseRecorder }
