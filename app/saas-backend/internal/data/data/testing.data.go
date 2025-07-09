package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/po"
	datarepos "github.com/go-micro-saas/saas-backend/app/saas-backend/internal/data/repo"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	idpkg "github.com/ikaiguang/go-srv-kit/kit/id"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type backendData struct {
	log *log.Helper
}

func NewBackendData(
	logger log.Logger,
) datarepos.BackendDataRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "test-service/data/data"))

	return &backendData{
		log: logHelper,
	}
}

func (s *backendData) HelloWorld(ctx context.Context, dataModel *po.HelloWorld) error {
	dataModel.Id, _ = idpkg.NextID()
	var err error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

func (s *backendData) QueryByID(ctx context.Context, id uint64) (dataModel *po.HelloWorld, isNotFound bool, err error) {
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return dataModel, isNotFound, err
}

func (s *backendData) Create(ctx context.Context, dataModel *po.HelloWorld) error {
	dataModel.Id, _ = idpkg.NextID()
	var err error
	if err != nil {
		if gormpkg.IsErrDuplicatedKey(err) {
			e := errorpkg.ErrorDuplicateKey("")
			err = errorpkg.Wrap(e, err)
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return err
	}
	return nil
}
