package service

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	servicev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/services/auth"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
)

// RegisterServices 注册服务
// @return Services 用于wire
// @return func() = cleanup 关闭资源
// @return error 错误
func RegisterServices(
	hs *http.Server, gs *grpc.Server,
	backendV1Service servicev1.SrvSaasBackendV1Server,
	backendAuthV1Service servicev1.SrvSaasBackendAuthV1Server,
) (cleanuputil.CleanupManager, error) {
	// 先进后出
	var cleanupManager = cleanuputil.NewCleanupManager()
	// grpc
	if gs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：GRPC: backendV1Service")
		servicev1.RegisterSrvSaasBackendV1Server(gs, backendV1Service)
		servicev1.RegisterSrvSaasBackendAuthV1Server(gs, backendAuthV1Service)

		//cleanupManager.Append(cleanup)
	}

	// http
	if hs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：HTTP: backendV1Service")
		servicev1.RegisterSrvSaasBackendV1HTTPServer(hs, backendV1Service)
		servicev1.RegisterSrvSaasBackendAuthV1HTTPServer(hs, backendAuthV1Service)

		// special
		//RegisterSpecialRouters(hs, homeService, websocketService)

		//cleanupManager.Append(cleanup)
	}

	return cleanupManager, nil
}

//func RegisterSpecialRouters(hs *http.Server, homeService *HomeService, websocketService *WebsocketService) {
//	// router
//	router := mux.NewRouter()
//
//	stdlog.Println("|*** REGISTER_ROUTER：Root(/)")
//	router.HandleFunc("/", homeService.Homepage)
//	hs.Handle("/", router)
//
//	stdlog.Println("|*** REGISTER_ROUTER：Websocket")
//	router.HandleFunc("/ws/v1/testdata/websocket", websocketService.TestWebsocket)
//
//	// router
//	hs.Handle("/ws/v1/websocket", router)
//}
