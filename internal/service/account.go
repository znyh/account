package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/znyh/account/internal/base"
	"github.com/znyh/account/internal/model"
	pb "github.com/znyh/proto/account"
)

// 注册账号
func (s *Service) guestRegister(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterRsp, err error) {

	var id = int64(0)
	if id, err = s.getPoolUID(ctx); err != nil {
		log.Error("guestRegister error (get UID from Pool err)")
		return
	}

	u := &pb.UserInfo{
		Nick:          fmt.Sprintf("Royal_%d", id),
		Password:      req.Password,
		Telephone:     req.Telephone,
		OS:            req.OS,
		Sex:           0,
		UserID:        id,
		CreateTime:    base.GetCurSecString(),
		LastLoginTime: base.GetCurSecString(),
	}
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

	err = s.dao.GuestRegister(ctx, u, m)
	if err != nil {
		log.Error("guestRegister error (RedisHMSet err),u:%+v,m:%+v", u, m)
		return
	}
	resp = &pb.RegisterRsp{
		Info: u,
	}
	log.Info("%s", u.String())
	return
}

//从ID池里获取一个用户id
func (s *Service) getPoolUID(ctx context.Context) (int64, error) {
	key := model.GetHAccountIdPoolKey()
	field := model.GetHAccountIdPoolField()

	id, err := s.dao.RedisHIncr(ctx, key, field, "1")
	if err != nil {
		log.Error("getPoolUID error(RedisHIncr error) key:%+v field:%+v", key, field)
		return -1, err
	}
	if id < 1000000 {
		id = 1000001
		err = s.dao.RedisHSet(ctx, key, field, "1000001")
	}

	return id, err
}

// 注册机器人账户
func (s *Service) robotRegister(ctx context.Context, req *pb.RobotRegisterReq) (resp *pb.RobotRegisterRsp, err error) {
	return
}

// 获取账户信息
func (s *Service) getGuest(ctx context.Context, req *pb.GetGuestReq) (resp *pb.GetGuestRsp, err error) {
	return
}

// 绑定手机号
func (s *Service) bindTel(ctx context.Context, req *pb.BindTelReq) (resp *empty.Empty, err error) {
	return
}

// 解绑手机
func (s *Service) unbindTel(ctx context.Context, req *pb.UnbindTelReq) (resp *empty.Empty, err error) {
	return
}

// 获取短信验证码
func (s *Service) getSmsToken(ctx context.Context, req *pb.GetSmsTokenReq) (resp *pb.GetSmsTokenRsp, err error) {
	return
}

// 校检短信验证
func (s *Service) verifySmsToken(ctx context.Context, req *pb.VerifySmsTokenReq) (resp *empty.Empty, err error) {
	return
}

// 修改验证码有效时间
func (s *Service) modifySmsToken(ctx context.Context, req *pb.ModifySmsTokenReq) (resp *empty.Empty, err error) {
	return
}

// 获取验证码信息
func (s *Service) getSmsTokenInfo(ctx context.Context, req *pb.GetSmsTokenInfoReq) (resp *pb.GetSmsTokenInfoRsp, err error) {
	return
}

// 获取绑定标志
func (s *Service) getBindFlag(ctx context.Context, req *pb.GetBindFlagReq) (resp *pb.GetBindFlagRsp, err error) {
	return
}

// 加载绑定信息
func (s *Service) loadBind(ctx context.Context, req *pb.LoadBindReq) (resp *pb.LoadBindRsp, err error) {
	return
}

// 设置用户属性集
func (s *Service) getUserInfo(ctx context.Context, req *pb.SetUserInfoReq) (resp *pb.SetUserInfoRsp, err error) {
	return
}

// 查询用户指定属性集
func (s *Service) loadUser(ctx context.Context, req *pb.LoadUserReq) (resp *pb.LoadUserRsp, err error) {
	return
}

// 批量查询用户指定属性集
func (s *Service) batchLoadUser(ctx context.Context, req *pb.BatchLoadUserReq) (resp *pb.BatchLoadUserRsp, err error) {
	return
}
