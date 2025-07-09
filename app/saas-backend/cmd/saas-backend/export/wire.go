//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/biz"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/data"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/service/service"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		//setuputil.GetRecommendDBConn,
		// data
		data.NewBackendData,
		// biz
		biz.NewBackendBiz,
		// service
		service.NewBackendV1Service,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
