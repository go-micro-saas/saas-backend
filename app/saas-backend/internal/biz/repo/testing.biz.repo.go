package bizrepos

import (
	"context"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/bo"
)

type BackendBizRepo interface {
	HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error)
}
