package wopan

import (
	"context"
	"fmt"
)

type ReqOpt struct {
	NoAccessToken bool
	Ctx           context.Context
}

func (w *WoClient) Request(channel string, key string, body interface{}, resp interface{}, opt ReqOpt) ([]byte, error) {
	req := w.NewRequest()
	req.SetHeaders(map[string]string{
		"Origin":  "https://pan.wo.cn",
		"Referer": "https://pan.wo.cn/",
	})
	if !opt.NoAccessToken {
		req.SetHeader("Accesstoken", w.accessToken)
	}
	if opt.Ctx != nil {
		req.SetContext(opt.Ctx)
	}
	header := calHeader(key, channel)
	var _resp Resp
	req.SetBody(Json{
		"header": header,
		"body":   body,
	}).SetResult(&resp)
	res, err := req.Post(fmt.Sprintf("https://panservice.mail.wo.cn/%s/dispatcher", channel))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return res.Body(), fmt.Errorf("request failed with status: %s", res.Status())
	}
	if _resp.Status != "200" {
		return res.Body(), fmt.Errorf("request failed with status: %s, msg: %s", _resp.Status, _resp.Msg)
	}
	if _resp.Rsp.RspCode != "0000" {
		return res.Body(), fmt.Errorf("request failed with rsp_code: %s,rep_desc: %s", _resp.Rsp.RspCode, _resp.Rsp.RspDesc)
	}
	if resp != nil {
		if encrypted, ok := _resp.Rsp.Data.(string); ok {
			strData, err := w.crypto.Decrypt(encrypted, channel)
			if err != nil {
				return res.Body(), err
			}
			err = w.jsonUnmarshalFunc([]byte(strData), resp)
			if err != nil {
				return res.Body(), err
			}
		} else {
			CopyStruct(_resp.Rsp.Data, resp)
		}
	}
	return res.Body(), nil
}

func (w *WoClient) RequestApiUser(key string, body interface{}, resp interface{}, opt ReqOpt) ([]byte, error) {
	return w.Request(APIUserChannel, key, body, resp, opt)
}

func (w *WoClient) RequestWoHome(key string, body interface{}, resp interface{}, opt ReqOpt) ([]byte, error) {
	return w.Request(WoHomeChannel, key, body, resp, opt)
}
