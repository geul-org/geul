//ff:type feature=resolve type=model
//ff:what kindDirs — the path segment of each Kind under proposal/
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// kindDirs names the Kind segment of a codec spec address (GEUL-ORG 직접 코드 코덱북).
var kindDirs = map[core.Kind]string{core.KindReference: "reference", core.KindValue: "value", core.KindEdge: "edge", core.KindMeta: "meta"}
