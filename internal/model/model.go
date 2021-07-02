package model

import (
	"fmt"

	pb "github.com/znyh/proto/account"
)

var (
	APPID      = "account-service"
	TimeLayout = "2006-01-02 15:04:05"

	_HAccountIdPoolKey   = "HAccount:Seq"
	_HAccountIdPoolField = "seq"

	_HAccountGuestKey = "HAccount:%d" // key:userID
)

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

func GetHAccountGuestKey(id int64) string {
	return fmt.Sprintf(_HAccountGuestKey, id)
}
func GetHAccountIdPoolKey() string {
	return _HAccountIdPoolKey
}
func GetHAccountIdPoolField() string {
	return _HAccountIdPoolField
}

type TagUserInfo struct {
	*pb.UserInfo
}

func (u *TagUserInfo) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"Nick":          u.Nick,
		"Password":      u.Password,
		"Telephone":     u.Telephone,
		"OS":            u.OS,
		"Sex":           u.Sex,
		"UserID":        u.UserID,
		"CreateTime":    u.CreateTime,
		"LastLoginTime": u.LastLoginTime,
	}
	return m
}
