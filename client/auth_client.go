/*
@Time : 2022/1/27 16:34
@Author : Hwdhy
@File : auth_client
@Software: GoLand
*/
package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc_project/pb"
	"time"
)

type AuthCLient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthCLient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthCLient{
		service:  service,
		username: username,
		password: password,
	}
}

func (client *AuthCLient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil
}
