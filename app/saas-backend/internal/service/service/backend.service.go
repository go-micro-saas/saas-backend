package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	resourcev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/resources"
	servicev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/services"
	bizrepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/repo"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/service/dto"
)

type backendV1Service struct {
	servicev1.UnimplementedSrvSaasBackendV1Server

	log        *log.Helper
	backendBiz bizrepos.BackendBizRepo
}

func NewBackendV1Service(logger log.Logger, testingBiz bizrepos.BackendBizRepo) servicev1.SrvSaasBackendV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/service"))
	return &backendV1Service{
		log:        logHelper,
		backendBiz: testingBiz,
	}
}

func (s *backendV1Service) Ping(ctx context.Context, req *resourcev1.PingReq) (*resourcev1.PingResp, error) {
	param := dto.BackendDto.ToBoHelloWorldParam(req)
	reply, err := s.backendBiz.HelloWorld(ctx, param)
	if err != nil {
		return nil, err
	}
	return &resourcev1.PingResp{
		Data: dto.BackendDto.ToPbPingRespData(reply.Message),
	}, nil
}
