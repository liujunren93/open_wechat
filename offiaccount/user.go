package officialAccounts

//
type userListRes struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

type userInfoRes struct {
	Subscribe int    `json:"subscribe"`
	OpenID    string `json:"openid"`
	NickName  string `json:"nickname"`
	UnionID   string `json:"unionid"`
	Avatar    string `json:"headimgurl"`
}

func GetUserList(next string) (res *userListRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/get"
	get, err := ToDoFuncGet(api, &res, "next_openid", next)
	err = ToDo.Do(get)
	return res, err
}

func GetUserInfo(openId string) (res *userInfoRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/info"
	get, err := ToDoFuncGet(api, &res, "openid", openId, "lang")
	err = ToDo.Do(get)
	return res, err

}
