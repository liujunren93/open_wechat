package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liujunren93/wechat/offiaccount/token"
	"github.com/liujunren93/wechat/utils/helper"
	"net/http"
)

type sendEntity struct {
	apiUrl      string `json:"-"`
	Touser      string `json:"touser"`
	TemplateID  string `json:"template_id"`
	Url         string `json:"url"`
	Miniprogram struct {
		AppID    string `json:"appid"`
		Pagepath string `json:"pagepath"`
	} `json:"miniprogram"`
	Data map[string]map[string]string
}

// New create a sendEntity
// apiUrl default https://api.weixin.qq.com/cgi-bin/template/del_private_template"
// Url 跳转地址
func New(apiUrl, templateID, Url string) *sendEntity {
	if apiUrl == "" {
		apiUrl = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	}
	return &sendEntity{
		apiUrl:     apiUrl,
		TemplateID: templateID,
		Url:        Url,
	}
}

func (e *sendEntity) JumpMiniprogram(appID, pagepath string) {
	e.Miniprogram.AppID = appID
	e.Miniprogram.Pagepath = pagepath
}

func (e *sendEntity) AddData(key, value, color string) {
	e.Data[key] = map[string]string{"value": value, "color": color}
}

func (e *sendEntity) Send(accessFunc token.AccessFunc, toUser string) (*sendRes, error) {
	if accessFunc == nil {
		return nil, errors.New("accessFunc cannot be nil")
	}
	if toUser == "" {
		return nil, errors.New("toUser cannot be nil")
	}

	s, err := accessFunc()

	if err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf("%s?access_token=%s", e.apiUrl, s)
	e.Touser = toUser
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	var header = make(http.Header)
	header.Add("Content-Type", "application/json")

	re, err := helper.HttpPost(apiUrl, header, marshal)

	var res sendRes
	err = json.Unmarshal(re, &res)
	if err != nil {
		return nil, nil
	}
	return &res, nil
}
