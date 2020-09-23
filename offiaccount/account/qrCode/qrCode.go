package qrCode

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/wechat/offiaccount/token"
	"github.com/liujunren93/wechat/utils/helper"
	"net/http"
	"time"
)

type ticket struct {
	options options
}

type option func(*options)

type TicketRes struct {
	Ticket        string        `json:"ticket"`
	ExpireSeconds time.Duration `json:"expire_seconds"`
	Url           string        `json:"url"`
}

//New
func New(opts ...option) *ticket {
	t := ticket{options: defaultOptions}
	for _, opt := range opts {
		opt(&t.options)
	}

	return &t
}

func (t *ticket) Create(accessFunc token.AccessFunc) (res *TicketRes, err error) {
	token, err := accessFunc()
	apiUrl := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token
	if err != nil {
		return nil, err
	}
	var header = make(http.Header)
	header.Add("Content-Type", "application/json")
	marshal, err := json.Marshal(t.options)
	if err != nil {
		return nil, err
	}
	post, err := helper.HttpPost(apiUrl, header, marshal)
	json.Unmarshal(post, &res)
	return res, err

}

func (t *TicketRes) QrCode() {
	apiUrl:="https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket="+t.Ticket
	get, err := helper.HttpGet(apiUrl)
	fmt.Println(string(get),err)
}
