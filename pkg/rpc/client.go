package rpc

import (
	"context"
	"fmt"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/app"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/setting"
	"google.golang.org/grpc"
	"time"
)

var GrpcClient UserServiceClient

// InitRPC 初始化RPC
func InitRPC() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", setting.DaoServerSetting.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %s", err)
	}

	GrpcClient = NewUserServiceClient(conn)
}

// Register 注册接口
func Register(requestID string, name string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(setting.AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.Register(ctx, &RegisterRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warnf("call Register failed, err:%v", err)
		return err
	}

	return nil
}

// Auth 认证接口
func Auth(requestID string, name string, password string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(setting.AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.Auth(ctx, &AuthRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warnf("call Auth failed, err:%v", err)
		return "", "", err
	}

	return rsp.Nickname, rsp.Image, nil
}

// GetProfile 查询用户的属性信息
func GetProfile(requestID string, name string) (info *app.ProfileInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(setting.AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetProfile(ctx, &GetProfileRequest{RequestID: requestID, Name: name})
	if err != nil {
		log.Warnf("call GetProfile failed, err:%v", err)
		return nil, err
	}

	return &app.ProfileInfo{Nickname: rsp.Nickname, ImageID: rsp.ImageID}, nil
}

// GetHeadImage 查询用户的头像信息
func GetHeadImage(requestID string, imageID string) (image []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(setting.AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetHeadImage(ctx, &GetHeadImageRequest{RequestID: requestID, ImageID: imageID})
	if err != nil {
		log.Warnf("call GetHeadImage failed, err:%v", err)
		return nil, err
	}

	return rsp.Image, nil
}

// EditProfile 编辑用户的属性信息
func EditProfile(requestID string, name string, nickname string, image []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(setting.AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.EditProfile(ctx, &EditProfileRequest{RequestID: requestID,
		Name: name, Nickname: nickname, Image: image})
	if err != nil {
		log.Warnf("call EditProfile failed, err:%v", err)
		return err
	}

	return nil
}
