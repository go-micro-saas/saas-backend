package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/repo"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/po"
	datarepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/repo"
	"time"
)

type backendBiz struct {
	log         *log.Helper
	backendData datarepos.BackendDataRepo
}

func NewBackendBiz(
	logger log.Logger,
	testingData datarepos.BackendDataRepo,
) bizrepos.BackendBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/biz"))

	return &backendBiz{
		log:         logHelper,
		backendData: testingData,
	}
}

func (s *backendBiz) HelloWorld(ctx context.Context, param *bo.HelloWorldParam) (*bo.HelloWorldReply, error) {
	dataModel := s.toHelloWorkModel(param)
	err := s.backendData.HelloWorld(ctx, dataModel)
	if err != nil {
		return nil, err
	}
	return &bo.HelloWorldReply{Message: dataModel.RequestMessage}, nil
}

func (s *backendBiz) toHelloWorkModel(param *bo.HelloWorldParam) *po.HelloWorld {
	res := &po.HelloWorld{
		Id:             0,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
		RequestMessage: param.Message,
	}
	return res
}
