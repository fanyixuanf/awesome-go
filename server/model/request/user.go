// Package request
// @Description:
package request

// RegisterStruct User register structure
type RegisterStruct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'Admin'"`
	HeaderImg   string `json:"headerImg" gorm:"default:xxx"`
	AuthorityId string `json:"authorityId" gorm:"default:1"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// RegisterAndLoginStruct User login structure
type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	HeaderImg string `json:"header_img"`
	UserProfile string `json:"user_profile"`
}
