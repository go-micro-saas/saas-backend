package schedulers

import (
	"github.com/go-kratos/kratos/v2/log"
	bizrepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/biz/repo"
)

type testingScheduler struct {
	log *log.Helper
}

func NewTestingScheduler(
	logger log.Logger,
) bizrepos.TestingSchedulerRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/biz/scheduler"))

	return &testingScheduler{
		log: logHelper,
	}
}
