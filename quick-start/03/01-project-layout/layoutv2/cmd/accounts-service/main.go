package main

import (
	v1 "github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/api/google/accounts/v1"
	"github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/internal/accounts/biz"
	"github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/internal/accounts/data"
	"github.com/qq1060656096/go-guide/quick-start/02/01-project-layout/layoutv2/internal/accounts/service"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main()  {
	lis, err := net.Listen("tcp", ":18022")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := log.Default()
	s := grpc.NewServer()
	loginService := service.NewLoginService(biz.NewLoginUseCase(data.NewLoginRepo(nil, logger), logger), logger)
	v1.RegisterLoginServer(s, loginService)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
