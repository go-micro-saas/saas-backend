package dto

import (
	resourcev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/resources"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/bo"
)

var (
	BackendDto backendDto
)

type backendDto struct{}

func (s *backendDto) ToBoHelloWorldParam(req *resourcev1.PingReq) *bo.HelloWorldParam {
	res := &bo.HelloWorldParam{
		Message: req.Message,
	}
	return res
}

func (s *backendDto) ToPbPingRespData(msg string) *resourcev1.PingRespData {
	res := &resourcev1.PingRespData{
		Message: msg,
	}

	return res
}
