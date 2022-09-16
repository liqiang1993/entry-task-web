package rpc

import (
	"context"
	pb "github.com/lucky-cheerful-man/phoenix_apis/protobuf3.pb/user_info_manage"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/server"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/util"

	"time"
)

var GrpcClient pb.UserService

func init() {
	GrpcClient = pb.NewUserService("phoenix_server", server.ReferClient())
}

// Register 注册接口
func Register(requestID string, name string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.Register(ctx, &pb.RegisterRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warn("call Register failed, err:%v", err)
		return err
	}

	return nil
}

// Auth 认证接口
func Auth(requestID string, name string, password string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.Auth(ctx, &pb.AuthRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warn("call Auth failed, err:%v", err)
		return "", "", err
	}

	return rsp.Nickname, rsp.Image, nil
}

// GetProfile 查询用户的属性信息
func GetProfile(requestID string, name string) (info *util.ProfileInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetProfile(ctx, &pb.GetProfileRequest{RequestID: requestID, Name: name})
	if err != nil {
		log.Warn("call GetProfile failed, err:%v", err)
		return nil, err
	}

	return &util.ProfileInfo{Nickname: rsp.Nickname, ImageID: rsp.ImageID}, nil
}

// GetHeadImage 查询用户的头像信息
func GetHeadImage(requestID string, imageID string) (image []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetHeadImage(ctx, &pb.GetHeadImageRequest{RequestID: requestID, ImageID: imageID})
	if err != nil {
		log.Warn("call GetHeadImage failed, err:%v", err)
		return nil, err
	}

	return rsp.Image, nil
}

// EditProfile 编辑用户的属性信息
func EditProfile(requestID string, name string, nickname string, image []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.EditProfile(ctx, &pb.EditProfileRequest{RequestID: requestID,
		Name: name, Nickname: nickname, Image: image})
	if err != nil {
		log.Warn("call EditProfile failed, err:%v", err)
		return err
	}

	return nil
}
