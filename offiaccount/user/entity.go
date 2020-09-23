package user

type (
	userListDataEntity struct {
		OpenID     []string `json:"openid"`
		NextOpenID string   `json:"next_openid"`
	}
	// userListEntity 公众号获取用户列表
	userListEntity struct {
		Total int                `json:"total"`
		Count int                `json:"count"`
		Data  userListDataEntity `json:"data"`
	}
)
