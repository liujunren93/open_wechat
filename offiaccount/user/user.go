package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liujunren93/wechat/offiaccount/token"
	"github.com/liujunren93/wechat/utils/helper"
)

type Next func(string) (*userListEntity, error)

//GetUserList 获取用户列表
// apiUrl default https://api.weixin.qq.com/cgi-bin/user/get
func GetUserList(accessFunc token.AccessFunc, apiUrl string) (Next, error) {
	if accessFunc == nil {
		return nil, errors.New("accessFunc cannot be nil")
	}
	if apiUrl == "" {
		apiUrl = "https://api.weixin.qq.com/cgi-bin/user/get"
	}
	return func(nextID string) (*userListEntity, error) {
		token, err := accessFunc()
		if err != nil {
			return nil, err
		}
		tmpUrl := fmt.Sprintf("%s?ACCESS_TOKEN=%s&next_openid=%s", apiUrl, token, nextID)
		get, err := helper.HttpGet(tmpUrl)
		if err != nil {
			return nil, err
		}
		var list userListEntity
		err = json.Unmarshal(get, &list)

		if err != nil {
			return nil, errors.New(string(get))
		}
		return &list, nil
	}, nil

}
