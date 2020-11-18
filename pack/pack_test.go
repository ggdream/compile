package pack

import "testing"


func TestZipPacker_Pack(t *testing.T) {
	z := ZipPacker{"tgz.go", "a.zip"}
	if err := z.Pack(); err != nil {
		panic(err)
	}
}

func TestTgzPacker_Pack(t *testing.T) {
	a := TgzPacker{"/home/moca/git/compile/conf/config.go", "a.tgz"}
	if err := a.Pack(); err != nil {
		panic(err)
	}
}
