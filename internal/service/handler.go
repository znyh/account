package service

import (
	"context"

	pb "github.com/znyh/proto/account"

	"github.com/golang/protobuf/ptypes/empty"
)

// 注册账号
func (s *Service) GuestRegister(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterRsp, err error) {
	return s.guestRegister(ctx, req)
}

// 注册机器人账户
func (s *Service) RobotRegister(ctx context.Context, req *pb.RobotRegisterReq) (resp *pb.RobotRegisterRsp, err error) {
	return
}

// 获取账户信息
func (s *Service) GetGuest(ctx context.Context, req *pb.GetGuestReq) (resp *pb.GetGuestRsp, err error) {
	return
}

// 绑定手机号
func (s *Service) BindTel(ctx context.Context, req *pb.BindTelReq) (resp *empty.Empty, err error) {
	return
}

// 解绑手机
func (s *Service) UnbindTel(ctx context.Context, req *pb.UnbindTelReq) (resp *empty.Empty, err error) {
	return
}

// 获取短信验证码
func (s *Service) GetSmsToken(ctx context.Context, req *pb.GetSmsTokenReq) (resp *pb.GetSmsTokenRsp, err error) {
	return
}

// 校检短信验证
func (s *Service) VerifySmsToken(ctx context.Context, req *pb.VerifySmsTokenReq) (resp *empty.Empty, err error) {
	return
}

// 修改验证码有效时间
func (s *Service) ModifySmsToken(ctx context.Context, req *pb.ModifySmsTokenReq) (resp *empty.Empty, err error) {
	return
}

// 获取验证码信息
func (s *Service) GetSmsTokenInfo(ctx context.Context, req *pb.GetSmsTokenInfoReq) (resp *pb.GetSmsTokenInfoRsp, err error) {
	return
}

// 获取绑定标志
func (s *Service) GetBindFlag(ctx context.Context, req *pb.GetBindFlagReq) (resp *pb.GetBindFlagRsp, err error) {
	return
}

// 加载绑定信息
func (s *Service) LoadBind(ctx context.Context, req *pb.LoadBindReq) (resp *pb.LoadBindRsp, err error) {
	return
}

// 设置用户属性集
func (s *Service) SetUserInfo(ctx context.Context, req *pb.SetUserInfoReq) (resp *pb.SetUserInfoRsp, err error) {
	return
}

// 查询用户指定属性集
func (s *Service) LoadUser(ctx context.Context, req *pb.LoadUserReq) (resp *pb.LoadUserRsp, err error) {
	return
}

// 批量查询用户指定属性集
func (s *Service) BatchLoadUser(ctx context.Context, req *pb.BatchLoadUserReq) (resp *pb.BatchLoadUserRsp, err error) {
	return
}
