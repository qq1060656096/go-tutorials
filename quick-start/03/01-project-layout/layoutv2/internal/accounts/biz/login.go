package biz

import (
	"context"
	"log"
)

type LoginAccount struct {
	AccountName string
	Password string
}

type LoginAccountRepo interface {
	Login(context.Context, *LoginAccount) error
}

type LoginUseCase struct {
	repo LoginAccountRepo
	log  *log.Logger
}

func NewLoginUseCase(repo LoginAccountRepo, logger *log.Logger) *LoginUseCase {
	return &LoginUseCase{
		repo: repo,
		log: logger,
	}
}
func (luc *LoginUseCase) AccountLogin(ctx context.Context, la *LoginAccount) error {
	return luc.repo.Login(ctx, la)
}
