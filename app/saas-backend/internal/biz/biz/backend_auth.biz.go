package biz

import (
	"context"
	resources "github.com/go-micro-saas/service-api/api/account-service/v1/resources"
	apiutil "github.com/go-micro-saas/service-api/util"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

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
