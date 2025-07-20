//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/biz"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/data"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/service/dto"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/service/service"
	accountapi "github.com/go-micro-saas/service-api/app/account-service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		//setuputil.GetRecommendDBConn,
		setuputil.GetServiceAPIManager,
		// data
		data.NewBackendData,
		// api
		dto.GetAccountV1ServiceNameForGRPC,
		accountapi.NewAccountV1GRPCClient,
		accountapi.NewUserAuthV1GRPCClient,
		//dto.GetAccountV1ServiceNameForHTTP,
		// accountapi.NewAccountV1HTTPClient,
		// biz
		biz.NewBackendBiz,
		// service
		service.NewBackendV1Service, service.NewBackendAuthV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
