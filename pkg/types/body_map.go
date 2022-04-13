package types

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/yunmoon/itp-request-client/pkg/util"
	"net/url"
	"sort"
	"strings"
)

type BodyMap map[string]interface{}

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) BodyMap {
	bm[key] = value
	return bm
}

func (bm BodyMap) SetBodyMap(key string, value func(bm BodyMap)) BodyMap {
	_bm := make(BodyMap)
	value(_bm)
	bm[key] = _bm
	return bm
}

// 获取参数，同 GetString()
func (bm BodyMap) Get(key string) string {
	return bm.GetString(key)
}

// 获取参数转换string
func (bm BodyMap) GetString(key string) string {
	if bm == nil {
		return ""
	}
	value, ok := bm[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

// 获取原始参数
func (bm BodyMap) GetInterface(key string) interface{} {
	if bm == nil {
		return nil
	}
	return bm[key]
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
}

// 置空BodyMap
func (bm BodyMap) Reset() {
	for k := range bm {
		delete(bm, k)
	}
}

func (bm BodyMap) JsonBody() (jb string) {
	bs, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(bm)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeURLParams() string {
	if bm == nil {
		return ""
	}
	var (
		str  string
		keys []string
	)
	for k := range bm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := bm.GetString(k); v != "" {
			encodeStr := url.QueryEscape(v)
			str = util.StrAppend(str, url.QueryEscape(k), "=", encodeStr, "&")
		}
	}
	if str == "" {
		return str
	}
	return str[:len(str)-1]
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == "" {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return errors.New(strings.Join(emptyKeys, ", ") + " : cannot be empty")
	}
	return nil
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
