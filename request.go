package wopan

import (
	"fmt"
)

func (w *WoClient) Request(channel string, key string, param, other Json, resp interface{}, opts ...RestyOption) ([]byte, error) {
	req := w.NewRequest()
	req.SetHeaders(map[string]string{
		"Origin":  "https://pan.wo.cn",
		"Referer": "https://pan.wo.cn/",
	})
	if w.accessToken != "" {
		req.SetHeader("Accesstoken", w.accessToken)

	}
	header := calHeader(channel, key)
	body, err := w.NewBody(channel, param, other)
	if err != nil {
		return nil, err
	}
	var _resp Resp
	req.SetBody(Req[interface{}]{
		Header: header,
		Body:   body,
	}).SetResult(&_resp)
	for _, opt := range opts {
		opt(req)
	}
	res, err := req.Post(fmt.Sprintf("%s%s/dispatcher", BaseURL, channel))
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

func (w *WoClient) RequestApiUser(key string, param, other Json, resp interface{}, opts ...RestyOption) ([]byte, error) {
	return w.Request(APIUserChannel, key, param, other, resp, opts...)
}

func (w *WoClient) RequestWoHome(key string, param, other Json, resp interface{}, opts ...RestyOption) ([]byte, error) {
	return w.Request(WoHomeChannel, key, param, other, resp, opts...)
}
