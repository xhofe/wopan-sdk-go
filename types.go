package wopan

type Json map[string]interface{}

type Header struct {
	Key     string `json:"key"`
	ResTime int64  `json:"resTime"`
	ReqSeq  int    `json:"reqSeq"`
	Channel string `json:"channel"`
	Sign    string `json:"sign"`
	Version string `json:"version"`
}

type Req[T any] struct {
	Header `json:"header"`
	Body   T `json:"body"`
}

type BodyWithParam struct {
	Param    string `json:"param"`
	ClientId string `json:"clientId"`
	Secret   bool   `json:"secret"`
}

type BodyWithKey struct {
	Param string `json:"param"`
	Key   bool   `json:"key"`
}

type BodyWithT[T any] struct {
	Param    T      `json:"param"`
	ClientId string `json:"clientId"`
}

type BodyWithParamNoClientId struct {
	Param  string `json:"param"`
	Secret bool   `json:"secret"`
}

type BodyWithQry struct {
	QryType  int    `json:"qryType"`
	ClientId string `json:"clientId"`
}

type ReqWithParam Req[BodyWithParam]
type ReqWithQry Req[BodyWithQry]
type ReqWithParamNoClientId Req[BodyWithParamNoClientId]
type ReqWithT[T any] Req[BodyWithT[T]]
type ReqWithKey Req[BodyWithKey]

type Resp struct {
	STATUS string `json:"STATUS"`
	MSG    string `json:"MSG"`
	LOGID  string `json:"LOGID"`
	RSP    struct {
		RSPCODE string      `json:"RSP_CODE"`
		RSPDESC string      `json:"RSP_DESC"`
		DATA    interface{} `json:"DATA"`
	} `json:"RSP"`
}
