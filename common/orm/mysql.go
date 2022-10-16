package orm

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"context"
	"shopping-cart/common/config"
	"shopping-cart/common/orm/logger"
)

var (
	reload      bool
	reloadMutex sync.Mutex
)

type dbKey struct{}

type Database struct {
	*gorm.DB
}

var DB *Database

func Setup(databaseConfig config.DatabaseConfig, debug bool) {
	var err error
	gormDB, err := gorm.Open(mysql.Open(databaseConfig.ConnectUrls.Master), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		Logger:                 logger.NewXZLogBridge(),
	})
	if err != nil {
		panic(fmt.Sprintf("orm open master node failed, err: %v", err))
	}
	dials := make([]gorm.Dialector, len(databaseConfig.ConnectUrls.Slaves))
	for _, slave := range databaseConfig.ConnectUrls.Slaves {
		dials = append(dials, mysql.Open(slave))
	}
	plugin := dbresolver.Register(dbresolver.Config{
		Replicas: dials,
	})

	err = gormDB.Use(plugin)
	if err != nil {
		panic(fmt.Sprintf("db use slave failed, err: %v", err))
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(fmt.Sprintf("get sql db failed, err: %v", err))
	}

	sqlDB.SetMaxIdleConns(databaseConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(databaseConfig.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(databaseConfig.LifeTime * int(time.Second)))

	DB = &Database{DB: gormDB}
	if debug {
		DB.DB = DB.DB.Debug()
	}
}

func (db *Database) Use(name string) *gorm.DB {
	if reload {
		reloadMutex.Lock()
		defer reloadMutex.Unlock()
	}

	return db.Clauses(dbresolver.Use(name))
}

func (db *Database) UseMaster() *gorm.DB {
	if reload {
		reloadMutex.Lock()
		defer reloadMutex.Unlock()
	}

	return db.Clauses(dbresolver.Write)
}

func NewContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func FromContext(ctx context.Context, name ...string) *gorm.DB {
	if reload {
		reloadMutex.Lock()
		defer reloadMutex.Unlock()
	}

	if len(name) > 0 {
		return DB.Use(name[0])
	}

	ins := DB.DB
	if db := ctx.Value(dbKey{}); db != nil {
		ins = db.(*gorm.DB)
	}

	return ins
}
