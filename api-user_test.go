package wopan_test

import (
	"testing"

	"github.com/Xhofe/wopan-sdk-go"
)

//func TestWoClient_RefreshToken(t *testing.T) {
//	w := wopan.DefaultWithRefreshToken("b34dbdbc-xxxx-46fd-96d2-ffd9516d87cc")
//	res, err := w.RefreshToken()
//	if err != nil {
//		t.Errorf("RefreshToken() error = %v", err)
//	} else {
//		t.Logf("RefreshToken() = %v", string(res))
//	}
//}

func TestWoClient_PcWebLogin(t *testing.T) {
	w := wopan.Default()
	_, err := w.PcWebLogin("", "")
	if err != nil {
		t.Logf("PcWebLogin() error = %v", err)
	}
}

func TestWoClient_AppQueryUser(t *testing.T) {
	w := wopan.DefaultWithAccessToken("91d4b946-xxxx-4909-bac1-d9914e45f2de")
	res, err := w.AppQueryUser()
	if err != nil {
		t.Errorf("AppQueryUser() error = %v", err)
	} else {
		t.Logf("AppQueryUser() = %+v", res)
	}
}
