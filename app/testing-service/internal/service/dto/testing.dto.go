package dto

import (
	resourcev1 "github.com/go-micro-saas/saas-backend/api/testing-service/v1/resources"
	"github.com/go-micro-saas/saas-backend/app/testing-service/internal/biz/bo"
)

var (
	TestingDto testingDto
)

type testingDto struct{}

func (s *testingDto) ToBoHelloWorldParam(req *resourcev1.TestReq) *bo.HelloWorldParam {
	res := &bo.HelloWorldParam{
		Message: req.Message,
	}
	return res
}

func (s *testingDto) ToPbTestRespData(msg string) *resourcev1.TestRespData {
	res := &resourcev1.TestRespData{
		Message: msg,
	}

	return res
}
