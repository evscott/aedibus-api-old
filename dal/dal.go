package dal

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/go-pg/pg/v9"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Info struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	Migrations string
}

type DAL struct {
	Log  *logger.StandardLogger
	DB   *pg.DB
	info *Info
}

func NewDAL(log *logger.StandardLogger, host, port, user, password, name, migrations string) *DAL {
	dal := &DAL{
		Log: log,
		info: &Info{
			Host:       host,
			Port:       port,
			User:       user,
			Password:   password,
			Name:       name,
			Migrations: migrations,
		},
	}

	dal.runMigrations()
	dal.setupGoPG()

	return dal
}

func (d *DAL) runMigrations() {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", d.info.Host, d.info.User, d.info.Password, d.info.Name)
	db, err := sql.Open(constants.DB_DRIVER, dbInfo)
	if err != nil {
		d.Log.ConfigError(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		d.Log.ConfigError(err)
	}
	m, err := migrate.NewWithDatabaseInstance(d.info.Migrations, d.info.Name, driver)
	if err != nil {
		d.Log.ConfigError(err)
	}
	if err := m.Up(); err != nil {
		d.Log.ConfigError(err)
	} else {
		log.Printf("Successfully ran migrations")
	}

	if err := db.Close(); err != nil {
		d.Log.ConfigError(err)
	}
}

func (d *DAL) setupGoPG() {
	d.DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", d.info.Host, d.info.Port),
		User:     d.info.User,
		Password: d.info.Password,
		Database: d.info.Name,
	})

	if _, err := d.DB.Exec("SELECT 1"); err != nil {
		d.Log.ConfigError(err)
	}
}
