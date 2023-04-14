package delivery

import (
	"fmt"
	"go-paybro/config"
	"go-paybro/controller"
	"go-paybro/manager"
	"log"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (a *AppServer) ver1() {
	ver1Routes := a.engine.Group("/ver1")
	a.customerController(ver1Routes)
}

func (a *AppServer) customerController(rg *gin.RouterGroup) {
	controller.NewCustomerController(rg, a.usecaseManager.CustomerUsecase())
}

func (a *AppServer) Run() {
	a.ver1()
	err := a.engine.Run(a.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfiguration()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
