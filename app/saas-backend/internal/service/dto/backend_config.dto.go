package dto

import (
	accountapi "github.com/go-micro-saas/service-api/app/account-service"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
)

func GetAccountV1ServiceNameForGRPC() []clientutil.ServiceName {
	return []clientutil.ServiceName{
		accountapi.AccountServiceGRPC,
	}
}

func GetAccountV1ServiceNameForHTTP() []clientutil.ServiceName {
	return []clientutil.ServiceName{
		accountapi.AccountServiceHTTP,
	}
}

func GetOrgV1ServiceNameForGRPC() []clientutil.ServiceName {
	return []clientutil.ServiceName{}
}

func GetOrgV1ServiceNameForHTTP() []clientutil.ServiceName {
	return []clientutil.ServiceName{}
}
