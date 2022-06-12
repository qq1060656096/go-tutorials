package data

import (
	"context"
	"database/sql"
	"github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/internal/accounts/biz"
	"log"
)

type loginRepo struct {
	db *sql.DB
	log  *log.Logger
}

func NewLoginRepo(db *sql.DB, logger *log.Logger) *loginRepo {
	return &loginRepo{
		db: db,
		log:  logger,
	}
}

func (lr *loginRepo) Login(ctx context.Context, la *biz.LoginAccount) error {
	return nil
}

