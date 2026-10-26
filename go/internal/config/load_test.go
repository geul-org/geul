//ff:func feature=config type=helper control=sequence
//ff:what TestLoad — flag over environment over file over default, GEULENV's source; a bad format and a broken file fail
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	conf := t.TempDir()
	file := filepath.Join(conf, "geul", "env")
	os.MkdirAll(filepath.Dir(file), 0o755)
	os.WriteFile(file, []byte("GEULDATA=/file\nGEULISSUER=me\nGEULREMOTE=https://x\n"), 0o644)
	env := map[string]string{"GEULISSUER": "you"}
	e := Env{Flags: map[string]string{"GEULDATA": "/flag"}, Getenv: func(k string) string { return env[k] }, Cwd: t.TempDir(), ConfigDir: conf}
	c, err := Load(e)
	if err != nil || c.Get("GEULDATA") != "/flag" || c.Get("GEULISSUER") != "you" || c.Source("GEULREMOTE") != "file" || c.Source("GEULFORMAT") != "default" || c.Source("GEULENV") != "default" || c.File != file {
		t.Fatalf("%+v %v", c, err)
	}
	env["GEULFORMAT"] = "xml"
	if _, err := Load(e); err == nil {
		t.Fatal("bad format: want error")
	}
	env["GEULFORMAT"] = ""
	os.WriteFile(file, []byte("junk\n"), 0o644)
	if _, err := Load(e); err == nil {
		t.Fatal("broken file: want error")
	}
}
