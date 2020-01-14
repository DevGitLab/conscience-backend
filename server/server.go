package server

import (
	"conscience-backend/bao"
	"conscience-backend/bao/meeting_chain"
	"conscience-backend/config"
	"conscience-backend/dao"
	"conscience-backend/db"
	"conscience-backend/http"
	"conscience-backend/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	http *http.Server
}

func New() *Server {
	s := new(Server)
	s.initialize()
	return s
}

func (s *Server) initialize() {
	config.Init()

	err := db.Init(config.GlobalConfig.DB)
	if err != nil {
		panic(err)
	}

	dao.Init(db.DB)

	api := bao.NewBlockchainAPI(config.GlobalConfig.BlockchainAPI)
	meeting_chain.Init(api)

	service.Init()

	s.http = http.New()
}

func (s *Server) Start() {
	go s.http.Start(config.GlobalConfig.Server.Addr)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for sig := range sigChan {
		fmt.Printf("get a signal %s", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			return
		case syscall.SIGHUP:
			//logger.Rotate(false)
		default:
			return
		}
	}
}
