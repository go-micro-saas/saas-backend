package backendapi

import (
	servicev1 "github.com/go-micro-saas/saas-backend/api/saas-backend/v1/services"
	_ "github.com/gorilla/websocket"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[servicev1.OperationSrvSaasBackendV1Ping] = middlewareutil.TransportServiceKindALL

	// auth
	whiteList[servicev1.OperationSrvSaasBackendAuthV1LoginByEmail] = middlewareutil.TransportServiceKindALL
	whiteList[servicev1.OperationSrvSaasBackendAuthV1LoginByPhone] = middlewareutil.TransportServiceKindALL

	return whiteList
}
