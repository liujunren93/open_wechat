package token

type accessTokenEntity struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateAt    int64
}
