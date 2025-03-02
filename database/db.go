package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goginmvc/globalconst"
	"goginmvc/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	if dbconfig.Password == "" {
		// if password is empty, use ''
		dbconfig.Password = "''"
	}

	dbStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.DBName)

	db, err := sqlx.Open("postgres", dbStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database. err: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	dbInstance = db

	return dbInstance, nil
}

func LoadDBConfig(env string) (*DBConfig, error) {
	if env == globalconst.PROD {
		return loadFromEnv()
	}

	// If development is enabled, load from local file
	return loadFromLocalFile()
}

func MockDB(db *sql.DB, databaseName string) {
	dbInstance = sqlx.NewDb(db, databaseName)
}

func loadFromEnv() (*DBConfig, error) {
	var dbConfig DBConfig
	keys := map[string]*string{
		"DB_USERNAME": &dbConfig.Username,
		"DB_PASSWORD": &dbConfig.Password,
		"DB_HOST":     &dbConfig.Host,
		"DB_PORT":     &dbConfig.Port,
		"DB_NAME":     &dbConfig.DBName,
	}
	for key, val := range keys {
		v, err := utils.LoadEnvVal(key)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s from environment variables. err: %w", key, err)
		}
		*val = v
	}

	return &dbConfig, nil
}

func loadFromLocalFile() (*DBConfig, error) {
	v := viper.New()
	v.SetConfigName(".env_dev")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

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
