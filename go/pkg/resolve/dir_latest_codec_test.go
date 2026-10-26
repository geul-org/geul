//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirLatestCodec — the latest version; missing, empty-list and tampered documents fail
package resolve

import (
	"bytes"
	"errors"
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestDirLatestCodec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x1018, 4))
	d.PublishCodec(codecDocOf(0x1018, 4, 2))
	if b, err := d.latestCodec(0x1018); err != nil || !bytes.Equal(b, codecDocOf(0x1018, 4, 2)) {
		t.Fatalf("latest: %v", err)
	}
	if _, err := d.latestCodec(0x1019); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
	k := codecKey(0x101A)
	d.write(listPath(k), listDocOf(0, VersionEntry{Version: 1, Deleted: true}))
	if _, err := d.latestCodec(0x101A); err == nil {
		t.Fatal("bad list: want error")
	}
	d.write(listPath(k), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0}, Doc: codecDocID(0x101A), Version: 1})}, nil, 0, 0))
	if _, err := d.latestCodec(0x101A); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no file: %v", err)
	}
	d.write(versionPath(k, 1), codecDocOf(0x101A))
	if _, err := d.latestCodec(0x101A); !errors.Is(err, ErrVerify) {
		t.Fatalf("tampered: %v", err)
	}
	d.write(listPath(k), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0}, Doc: codecDocID(0x101A), Version: 1}), core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 2, Issuer: []core.Word{0}, Doc: codecDocID(0x101A), Version: 1})}, nil, 0, 0))
	if _, err := d.latestCodec(0x101A); !errors.Is(err, ErrNotFound) {
		t.Fatalf("deleted: %v", err)
	}
}
