package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/phamtai97/go-utils/utils/logger"
	"go.uber.org/zap"
)

// MySQLConfig contains config of mysql.
//
// PoolSize is the maximum number of open connections to the database.
//
// MaxIdleConns is the maximum number of connections in the idle connection pool. 0 <= MaxIdleConns <= PoolSize
//
// ConnMaxLifetimeInMs is the maximum amount of time a connection may be reused.
// Expired connections may be closed lazily before reuse.
// If d <= 0, connections are reused forever.
type MySQLConfig struct {
	User                      string
	Password                  string
	Host                      string
	Port                      int
	DBName                    string
	PoolName                  string
	PoolSize                  int
	MaxIdleConns              int
	ConnMaxLifetimeInMs       int64
	ReadTimeoutInMs           int64
	WriteTimeoutInMs          int64
	DialConnectionTimeoutInMs int64
}

// MySQLImpl implement database interface
type MySQLImpl struct {
	db     *sqlx.DB
	config MySQLConfig
}

// NewMySQLImpl create instance mysql
func NewMySQLImpl(config MySQLConfig) (Database, error) {
	mysql := &MySQLImpl{
		config: config,
	}

	if err := mysql.Connect(); err != nil {
		return nil, err
	}

	return mysql, nil
}

// Connect to mysql server
func (mysql *MySQLImpl) Connect() error {
	db, err := sqlx.Connect("mysql", mysql.getDataSourceName())
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(mysql.config.PoolSize)
	db.SetMaxIdleConns(mysql.config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(mysql.config.ConnMaxLifetimeInMs) * time.Millisecond)

	mysql.db = db

	logger.Info("Connect Database success", zap.String("Pool name", mysql.config.PoolName))
	return nil
}

// Disconnect closes the database and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (mysql *MySQLImpl) Disconnect() error {
	return mysql.db.Close()
}

// GetConnection get db connection from pool
func (mysql *MySQLImpl) GetConnection() interface{} {
	return mysql.db
}

func (mysql *MySQLImpl) getDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms",
		mysql.config.User, mysql.config.Password, mysql.config.Host, mysql.config.Port, mysql.config.DBName,
		mysql.config.DialConnectionTimeoutInMs, mysql.config.ReadTimeoutInMs, mysql.config.WriteTimeoutInMs)
}
