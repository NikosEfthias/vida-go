package storage

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"testing"

	"gitlab.mugsoft.io/vida/go-api/config"
	"gitlab.mugsoft.io/vida/go-api/helpers/drivers/files/fs"
)

func TestService_public_files(t *testing.T) { //{{{
	//invalid filename{{{
	_, _, err := Service_public_files("non existing file")
	if err != fs.E_NOT_FOUND {
		t.Fatal("expected not found found", err)
	}
	//valid filename
	//create an example image{{{
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	var a color.Color = color.Gray{Y: 125}
	for i := 20; i <= 80; i++ {
		img.Set(i, 50, a)
	}
	os.MkdirAll(config.Get("PUBLIC_FILES_PATH"), 0777)
	file, err := os.OpenFile(config.Get("PUBLIC_FILES_PATH")+"test.png", os.O_CREATE|os.O_WRONLY, 0777)
	defer os.RemoveAll(config.Get("PUBLIC_FILES_PATH"))
	__fail_on_err(err, t)
	defer file.Close()
	err = png.Encode(file, img)
	__fail_on_err(err, t) //}}}
	_, ctype, err := Service_public_files("test.png")
	__fail_on_err(err, t)
	if strings.ToLower(ctype) != "image/png" {
		t.Fatal("invalid mime type expected png found ", ctype)
	}
	// }}}
} //}}}
func __fail_on_err(err error, t *testing.T) { //{{{
	if nil != err {
		t.Fatal(err)
	}
} //}}}
