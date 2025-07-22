package model

import (
	"context"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func init() {
	viper.GetViper()
	sync.OnceFunc(func() {
		New(context.TODO())
	})
}

type Model struct {
	gorm.Model
}

// New 获取全局db连接
func New(ctx context.Context) *gorm.DB {
	if globalDB == nil {
		db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{
			SkipDefaultTransaction:                   false,
			DefaultTransactionTimeout:                0,
			NamingStrategy:                           nil,
			FullSaveAssociations:                     false,
			Logger:                                   nil,
			NowFunc:                                  nil,
			DryRun:                                   false,
			PrepareStmt:                              false,
			PrepareStmtMaxSize:                       0,
			PrepareStmtTTL:                           0,
			DisableAutomaticPing:                     false,
			DisableForeignKeyConstraintWhenMigrating: false,
			IgnoreRelationshipsWhenMigrating:         false,
			DisableNestedTransaction:                 false,
			AllowGlobalUpdate:                        false,
			QueryFields:                              false,
			CreateBatchSize:                          0,
			TranslateError:                           false,
			PropagateUnscoped:                        false,
			ClauseBuilders:                           nil,
			ConnPool:                                 nil,
			Dialector:                                nil,
			Plugins:                                  nil,
		})
		if err != nil {
			panic(err)
		}
		globalDB = db
	}

	return globalDB
}
