package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	servicev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/services/auth"
	bizrepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/repo"
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
)

type backendAuthV1Service struct {
	servicev1.UnimplementedSrvSaasBackendAuthV1Server

	log        *log.Helper
	backendBiz bizrepos.BackendBizRepo
}

func NewBackendAuthV1Service(logger log.Logger, testingBiz bizrepos.BackendBizRepo) servicev1.SrvSaasBackendAuthV1Server {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/service/backend_auth"))
	return &backendAuthV1Service{
		log:        logHelper,
		backendBiz: testingBiz,
	}
}

func (s *backendAuthV1Service) LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	return s.backendBiz.LoginByEmail(ctx, req)
}

func (s *backendAuthV1Service) LoginByPhone(ctx context.Context, req *resources.LoginByPhoneReq) (*resources.LoginResp, error) {
	return s.backendBiz.LoginByPhone(ctx, req)
}
