package wopan

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func calHeader(key, channel string) Json {
	resTime := time.Now().UnixMilli()
	reqSeq := rand.Int31n(8999) + 1e5
	version := ""
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%s%d%d%s%s", key, resTime, reqSeq, channel, version)))
	sign := hex.EncodeToString(m.Sum(nil))
	return Json{
		"key":     key,
		"resTime": resTime,
		"reqSeq":  reqSeq,
		"channel": channel,
		"version": version,
		"sign":    sign,
	}
}
