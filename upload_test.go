package wopan_test

import (
	"os"
	"testing"

	"github.com/Xhofe/wopan-sdk-go"
)

func TestWoClient_Upload2C(t *testing.T) {
	w := wopan.DefaultWithAccessToken("91d4b946-xxxx-4909-bac1-d9914e45f2de")
	f, err := os.Open("./upload.go")
	if err != nil {
		t.Errorf("Open file error = %v", err)
	}
	fStat, err := f.Stat()
	if err != nil {
		t.Errorf("Get file stat error = %v", err)
	}
	fid, err := w.Upload2C(wopan.SpaceTypeFamily, wopan.Upload2CFile{
		Name:        fStat.Name(),
		Size:        fStat.Size(),
		Content:     f,
		ContentType: "text/plain",
	}, "37135bfb5178467b857129a00931e24a", "9263097", wopan.Upload2COption{})
	if err != nil {
		t.Errorf("Upload2C() error = %v", err)
	} else {
		t.Logf("Upload2C() = %+v", fid)
	}
}
