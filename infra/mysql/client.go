package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

func NewClient(ctx context.Context, cfg *Config) (*sql.DB, error) {
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	c := mysql.Config{
		User:            cfg.Username,
		Passwd:          cfg.Password,
		Net:             "tcp",
		Addr:            fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DBName:          cfg.DatabaseName,
		Loc:             loc,
		MultiStatements: true,
		ParseTime:       true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
