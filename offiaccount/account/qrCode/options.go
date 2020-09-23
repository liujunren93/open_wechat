package qrCode

import (
	"github.com/liujunren93/wechat/utils/helper"
	"time"
)

type options struct {
	ExpireSeconds time.Duration `json:"expire_seconds"`
	ActionName    string        `json:"action_name"`
	ActionInfo    interface{}   `json:"action_info"`
	SceneID       uint32        `json:"scene_id"`
	SceneStr      string        `json:"scene_str"`
}

var defaultOptions = options{
	ExpireSeconds: 3600,
	ActionName:    "QR_STR_SCENE",
	ActionInfo:    nil,
	SceneID:       1,
	SceneStr:      helper.RandString(10),
}

//WithExpire
//max：2592000
func WithExpire(expire time.Duration) option {
	if expire > 2592000 {
		expire = 2592000
	}
	return func(o *options) {
		o.ExpireSeconds = expire
	}
}

//WithActionName
// 1 临时
//2 永久
func WithActionName(i int8) option {
	actionName := "QR_STR_SCENE"
	if i == 2 {
		actionName = "QR_LIMIT_SCENE"
	}
	return func(o *options) {
		o.ActionName = actionName
	}
}

//WithActionInfo 二维码详细信息
func WithActionInfo(info interface{}) option {
	return func(o *options) {
		o.ActionInfo = info
	}
}

//WithSceneID
//场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
func WithSceneID(sceneID uint32) option {
	return func(o *options) {
		o.SceneID = sceneID
	}
}
//WithSceneStr
//场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
func WithSceneStr(sceneStr string) option {
	return func(o *options) {
		o.SceneStr = sceneStr
	}
}
