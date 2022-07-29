package service

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
)

type DemoService struct{}

var (
	demoService     *DemoService
	demoServiceOnce sync.Once
)

func GetDemoService() *DemoService {
	demoServiceOnce.Do(func() {
		demoService = &DemoService{}
	})
	return demoService
}

func (ins *DemoService) Ping(c *gin.Context) (*bo.PingResponse, error) {
	pong, err := dal.GetDemoDal().Ping(c)
	if err != nil {
		return nil, err
	}
	return &bo.PingResponse{
		Pong: pong,
	}, nil
}

func (ins *DemoService) Hello(c *gin.Context, req *bo.HelloRequest) (*bo.HelloResponse, error) {
	if req.Username == "" {
		return nil, errors.New(common.USERINPUTERRORMSG)
	}
	hello, err := dal.GetDemoDal().Hello(c)
	if err != nil {
		return nil, err
	}
	demoString := hello.Hello + " " + req.Username + " from " + hello.Author
	return &bo.HelloResponse{
		Hello: demoString,
	}, nil
}
