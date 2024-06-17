package managers

import (
	"database/sql"
	"fmt"
	"restfull-api/m/v2/config"
)

type AppManager interface {
	Conn() *sql.DB
}

type appManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *appManager) Conn() *sql.DB {
	return i.db
}

func (i *appManager) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s db_name=%s sslmode=disable", i.cfg.DB.Host, i.cfg.DB.Port, i.cfg.DB.User, i.cfg.DB.Password, i.cfg.DB.Name)

	db, err := sql.Open(i.cfg.DB.Driver, dsn)
	if err != nil {
		return fmt.Errorf("Failed connection to database %v", err.Error())
	}

	i.db = db
	return nil
}

func Application(cfg *config.Config) (AppManager, error) {
	conn := &appManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}
