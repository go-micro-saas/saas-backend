package bizrepos

import (
	"context"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/bo"
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
)

type BackendBizRepo interface {
	HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error)

	LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq) (*resources.LoginResp, error)
	LoginByPhone(ctx context.Context, req *resources.LoginByPhoneReq) (*resources.LoginResp, error)
}
