package fs

import (
	"os"
	"testing"

	"gitlab.mugsoft.io/vida/go-api/config"
)

func TestPut(t *testing.T) { //{{{
	os.MkdirAll(config.Get("PUBLIC_FILES_PATH"), 0777)
	defer os.RemoveAll(config.Get("PUBLIC_FILES_PATH"))
	//empty string must fail//{{{
	err := Put("", nil)
	if E_INVALID_PATH != err {
		t.Fatalf("expected error E_INVALID_PATH found %v", err)
	} //}}}
	//valid name invalid file must fail{{{
	err = Put("test.bin", nil)
	if E_INVALID_DATA != err {
		t.Fatalf("expected error E_INVALID_DATA found %v", err)
	} //}}}
	//successful case{{{
	bin := []byte{1, 2, 3, 4, 5}
	err = Put("test.bin", bin)
	if nil != err {
		t.Fatalf("expected error to be nil found %v", err)
	}
	f, err := os.Open(config.Get("PUBLIC_FILES_PATH") + "test.bin")
	__fail_error(err, t)
	data := make([]byte, 20)
	n, err := f.Read(data)
	f.Close()
	__fail_error(err, t)
	if n != len(bin) {
		t.Fatal("data changed", n)
	}
	for i := range data[:n] {
		if bin[i] != data[i] {
			t.Fatal("data changed")
		}
	} //}}}

} //}}}
func TestGet(t *testing.T) { //{{{
	os.MkdirAll(config.Get("PUBLIC_FILES_PATH"), 0777)
	defer os.RemoveAll(config.Get("PUBLIC_FILES_PATH"))
	bin := []byte{1, 2, 3, 4}
	Put("test.bin", bin)
	//empty string must fail{{{
	_, err := Get("")
	if E_INVALID_PATH != err {
		t.Fatalf("expected error E_INVALID_PATH found %v", err)
	} //}}}
	//missing file{{{
	_, err = Get("missing.file")
	if err != E_NOT_FOUND {
		t.Fatal("expected:", E_NOT_FOUND, "\nfound:", err)
	} //}}}
	//actual file {{{
	d, err := Get("test.bin")
	__fail_error(err, t)
	if len(d) != len(bin) {
		t.Fatal("data changed")
	}
	for i := range d {
		if bin[i] != d[i] {
			t.Fatal("data changed")
		}
	} //}}}
} //}}}
func __fail_error(err error, t *testing.T) { //{{{
	if nil != err {
		t.Fatal(err)
	}
} //}}}
