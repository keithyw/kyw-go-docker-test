package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/keithyw/kyw-go-docker-test/conf"
)

type MysqlDB struct {
	Config *conf.Config
	DB *sql.DB
}

func NewDatabase(config *conf.Config) (*MysqlDB) {
	mysqlConfig := mysql.Config{
		User: config.MysqlUser,
		Passwd: config.MysqlPass,
		Net: "tcp",
		Addr: config.MysqlHost,
		DBName: config.MysqlDBName,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	// defer db.Close()
	return &MysqlDB{
		config, 
		db, 
	}
}