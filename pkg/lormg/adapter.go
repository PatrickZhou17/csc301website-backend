package lormg

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shopping-cart/common/orm"
	"shopping-cart/pkg/log"
)

// FromContext to get orm.DB
func FromContext(ctx context.Context) *gorm.DB {
	tx := orm.FromContext(ctx)
	tx.Logger = NewOrmLogger(ctx)
	return tx
}

// NewOrmLogger -
func NewOrmLogger(ctx context.Context) logger.Interface {
	olw := &ormLoggerWriter{
		ctx: ctx,
	}
	return logger.New(olw, logger.Config{}).LogMode(logger.Info)
}

type ormLoggerWriter struct {
	ctx context.Context
}

func (olw *ormLoggerWriter) Printf(format string, args ...interface{}) {
	log.GetLogger(olw.ctx, "gorm").Infof(format, args...)
}
