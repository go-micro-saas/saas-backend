package serviceexporter

import (
	backendapi "github.com/go-micro-saas/saas-backend/api/saas-backend"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/conf"
	"github.com/go-micro-saas/saas-backend/app/testing-service/cmd/database-migration/migrate"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
	return []map[string]middlewareutil.TransportServiceKind{
		backendapi.GetAuthWhiteList(),
	}
}

func ExportServices(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
	hs, err := serverManager.GetHTTPServer()
	if err != nil {
		return nil, err
	}
	gs, err := serverManager.GetGRPCServer()
	if err != nil {
		return nil, err
	}
	return exportServices(launcherManager, hs, gs)
	//return serverutil.MergeCleanup(exportServices(launcherManager, hs, gs))
}

func ExportServicesWithDatabaseMigration(launcherManager setuputil.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
	settingConfig := launcherManager.GetConfig().GetSetting()

	if settingConfig.GetEnableMigrateDb() {
		dbConn, err := setuputil.GetRecommendDBConn(launcherManager)
		if err != nil {
			return nil, err
		}
		logger, err := setuputil.GetLogger(launcherManager)
		if err != nil {
			return nil, err
		}
		dbmigrate.Run(dbConn, dbutil.WithLogger(logger))
	}
	return ExportServices(launcherManager, serverManager)
}

func ExportDatabaseMigration() []dbutil.MigrationFunc {
	return []dbutil.MigrationFunc{
		dbmigrate.Run,
	}
}
