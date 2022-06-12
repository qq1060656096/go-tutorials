package service

import (
	"context"
	v1 "github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/api/google/accounts/v1"
	"github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/internal/accounts/biz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// UserService is a user service.
type LoginService struct {
	v1.UnimplementedLoginServer
	luc  *biz.LoginUseCase
	log *log.Logger
}

// NewLoginService new a login service.
func NewLoginService(uc *biz.LoginUseCase, log *log.Logger) *LoginService {
	return &LoginService{luc: uc, log: log}
}


func (s *LoginService) AccountLogin(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	la := &biz.LoginAccount{
		AccountName: in.AccountName,
		Password: in.Password,
	}
	err := s.luc.AccountLogin(ctx, la)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}
	userResponse := &v1.LoginResponse{
		AccountName: in.AccountName,
	}
	return userResponse, nil
}
