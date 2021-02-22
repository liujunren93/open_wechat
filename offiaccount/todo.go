package officialAccounts

import (
	"blal/config"
	"blal/pkg/weChat"
	"blal/utils"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"
)

type todo struct {
	retry int32
	token *token
	mutex sync.RWMutex
}

var ToDo *todo

func init() {
	ToDo = new(todo)
}

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateAt    int64
}

type toDoFuncT func(token string) ([]byte, error)

func ToDoFuncGet(api string, res interface{}, kv ...string) (toDoFuncT, error) {
	var val = make(url.Values)
	if len(kv)%2 != 0 {
		return nil, errors.New("KV has to be even ")
	}
	for i := 0; i < len(kv)/2; i += 2 {
		val.Add(kv[i], kv[i+1])
	}
	return func(token string) ([]byte, error) {
		val.Add("access_token", token)
		api := api + "?" + val.Encode()
		re, err := utils.HttpGet(api)
		if err != nil {
			return re, err
		}
		err = json.Unmarshal(re, &res)

		return re, err
	}, nil
}

func (t *todo) Do(f toDoFuncT) error {
	errRes := new(weChat.ErrorRes)
	t.getToken()
	bytes, err := f(t.token.AccessToken)
	if err != nil {
		return err
	}
	if strings.Index(string(bytes), "errcode") > 0 {
		_ = json.Unmarshal(bytes, &errRes)
		if errRes.ErrorCode == 40001 {
			if t.retry >= 3 {
				return errRes
			}
			t.retry++
			return t.Do(f)
		}
		return errRes
	}

	return nil

}

func (t *todo) getToken() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	apiUrl := "https://api.weixin.qq.com/cgi-bin/token"
	if t.token == nil || t.retry > 0 || t.token.CreateAt == 0 || (time.Now().Unix()-t.token.CreateAt) >= 7100 {
		url := apiUrl + "?grant_type=client_credential&AppId=" + config.SysConf.WeChat.OfficialAccounts.AppId + "&secret=" + config.SysConf.WeChat.OfficialAccounts.AppSecret
		get, _ := utils.HttpGet(url)
		json.Unmarshal(get, &t.token)
		t.token.CreateAt = time.Now().Unix()
	}

}
