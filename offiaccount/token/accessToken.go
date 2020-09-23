package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liujunren93/wechat/utils/helper"
	"time"
)

type AccessFunc func() (string, error)

//AccessToken get accessToken
// apiUrl default : https://api.weixin.qq.com/cgi-bin/token
// grantType default : client_credential
func AccessToken(grantType, appID, secret, apiUrl string) AccessFunc {
	var atoken = new(accessTokenEntity)
	return func() (string, error) {

		if atoken.AccessToken == "" || (time.Now().Local().Unix()- atoken.CreateAt)>=7100 {
			var err error
			atoken, err = accessToken(grantType, appID, secret, apiUrl)
			if err != nil {
				return "", err
			}
		}
		return atoken.AccessToken, nil
	}
}

// accessToken get accessToken
func accessToken(grantType, appID, secret, apiUrl string) (*accessTokenEntity, error) {
	var atoken accessTokenEntity
	if apiUrl == "" {
		apiUrl = "https://api.weixin.qq.com/cgi-bin/token"
	}
	if grantType == "" {
		grantType = "client_credential"
	}
	apiUrl = fmt.Sprintf(apiUrl+"?grant_type=%s&appid=%s&secret=%s", grantType, appID, secret)
	get, err := helper.HttpGet(apiUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(get, &atoken)
	if err != nil {

		return nil, errors.New(string(get))
	}

	atoken.CreateAt = time.Now().Local().Unix()
	time.AfterFunc(time.Second*7100, func() {
		atoken.AccessToken = ""
	})
	if atoken.AccessToken=="" {
		return nil, errors.New(string(get))
	}
	return &atoken, nil

}
