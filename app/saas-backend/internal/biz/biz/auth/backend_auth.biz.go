package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/bo"
	bizrepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/repo"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/po"
	datarepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/repo"
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	accountservicev1 "github.com/go-micro-saas/service-api/api/account-service/v1/services"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"

	"time"
)

type backendBiz struct {
	log *log.Helper

	backendData datarepos.BackendDataRepo

	userAuthV1Client accountservicev1.SrvUserAuthV1Client
	accountV1Client  accountservicev1.SrvAccountV1Client
}

func NewBackendBiz(
	logger log.Logger,
	testingData datarepos.BackendDataRepo,
	userAuthV1Client accountservicev1.SrvUserAuthV1Client,
	accountV1Client accountservicev1.SrvAccountV1Client,
) bizrepos.BackendBizRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "saas-backend/biz/biz"))

	return &backendBiz{
		log:              logHelper,
		backendData:      testingData,
		userAuthV1Client: userAuthV1Client,
		accountV1Client:  accountV1Client,
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

func (s *backendBiz) LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	resp, err := s.userAuthV1Client.LoginByEmail(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp, nil
}

func (s *backendBiz) LoginByPhone(ctx context.Context, req *resources.LoginByPhoneReq) (*resources.LoginResp, error) {
	resp, err := s.userAuthV1Client.LoginByPhone(ctx, req)
	if e := apiutil.CheckAPIResponse(resp, err); e != nil {
		return nil, errorpkg.WithStack(e)
	}
	return resp, nil
}
