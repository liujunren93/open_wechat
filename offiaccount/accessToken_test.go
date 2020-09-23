package offiaccount

import (
	"fmt"
	"github.com/liujunren93/wechat/offiaccount/account/qrCode"
	"github.com/liujunren93/wechat/offiaccount/msg/template"
	token2 "github.com/liujunren93/wechat/offiaccount/token"
	"github.com/liujunren93/wechat/offiaccount/user"
	"testing"
)

func TestAccessToken(t *testing.T) {
	token := token2.AccessToken("", "wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", "")

	next, err := user.GetUserList(token, "")
	fmt.Println(err)
	entity, err:= next("")
	fmt.Println(entity,err)

	tp := template.New("", "q1aLDc36P6_5O5wy7q6xUc_La4ngYXQKizp4QNJSKos", "")
	for _, openID := range entity.Data.OpenID {
		send, err := tp.Send(token, openID)
		fmt.Println(send,err)
	}
}

func TestQrCode(t *testing.T)  {

	token := token2.AccessToken("", "wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", "")
	ticket := qrCode.New(qrCode.WithActionName(2))
	create, err := ticket.Create(token)
	fmt.Println("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket="+create.Ticket)
	t.Log(create, err)
}