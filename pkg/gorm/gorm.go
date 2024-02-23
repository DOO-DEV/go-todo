package gorm

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysqlGorm(db *sql.DB, debug bool) (*gorm.DB, error) {
	gormDb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect to gorm", zap.Error(err))
		return nil, err
	}

	if debug {
		gormDb.Logger = gormDb.Logger.LogMode(logger.Info)
	} else {
		gormDb.Logger = gormDb.Logger.LogMode(logger.Silent)
	}

	return gormDb, nil
}
