package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goginmvc/globalconst"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var dbInstance *sqlx.DB

type DBConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DBName   string `toml:"dbname"`
}

func InitDB(env string) (*sqlx.DB, error) {
	// load db config
	dbconfig, err := LoadDBConfig(env)
	if err != nil {
		return nil, fmt.Errorf("failed to load database config. err: %w", err)
	}
	// init db
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.DBName)
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database. err: %w", err)
	}

	dbInstance = db

	return dbInstance, nil
}

func LoadDBConfig(env string) (*DBConfig, error) {
	v := viper.New()
	v.SetConfigName(".env_dev")
	if env == globalconst.PROD {
		v.SetConfigName(".env_prod")
	}
	v.SetConfigType("toml")
	v.AddConfigPath("../")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read database config. err: %w", err)
	}

	b, err := json.Marshal(v.Get("database"))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal database config. err: %w", err)
	}

	dbConfig := new(DBConfig)
	if err := json.Unmarshal(b, &dbConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal database config. err: %w, config: %s", err, string(b))
	}

	return dbConfig, nil
}

func MockDB(db *sql.DB, databaseName string) {
	dbInstance = sqlx.NewDb(db, databaseName)
}
