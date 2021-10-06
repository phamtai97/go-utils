// Package database provides an interface Database that can be implemented by different data sources.
//
// Example Usage
//
// The following is a complete example using mysql data source of database package
// 	import (
// 		"database/sql"
//	
// 		"github.com/jmoiron/sqlx"
// 		"github.com/phamtai97/go-utils/utils/datetime"
// 		database "github.com/phamtai97/go-utils/utils/db"
// 		"github.com/phamtai97/go-utils/utils/logger"
// 		"go.uber.org/zap"
// 	)
//	
// 	// AccountDTO data transfer object
// 	type AccountDTO struct {
// 		ID          int64  `db:"id"`
// 		Username    string `db:"username"`
// 		Password    string `db:"password"`
// 		Email       string `db:"email"`
// 		Status      int    `db:"status"`
// 		Role        string `db:"role"`
// 		CreatedTime int64  `db:"created_time"`
// 		UpdatedTime int64  `db:"updated_time"`
// 	}
//	
// 	func main() {
// 		logger.InitProduction("")
//	
// 		config := database.MySQLConfig{
// 			User:                      "dbgtest",
// 			Password:                  "abc@123",
// 			Host:                      "10.30.17.173",
// 			Port:                      4000,
// 			DBName:                    "go_admin",
// 			PoolName:                  "account_da",
// 			PoolSize:                  10,
// 			MaxIdleConns:              2,
// 			ConnMaxLifetimeInMs:       10000,
// 			ReadTimeoutInMs:           3000,
// 			WriteTimeoutInMs:          3000,
// 			DialConnectionTimeoutInMs: 3000,
// 		}
//	
// 		mysql, err := database.NewMySQLImpl(config)
// 		if err != nil {
// 			logger.Fatal("Failed to create mysql", zap.Error(err))
// 		}
// 		defer mysql.Disconnect()
//	
// 		db := mysql.GetConnection().(*sqlx.DB)
//	
// 		//
// 		// Insert new account into DB
// 		newAccount := AccountDTO{
// 			Username:    "AJPham",
// 			Password:    "123@ajpham",
// 			Email:       "go-util@gmail.com",
// 			Status:      1,
// 			Role:        "admin",
// 			CreatedTime: datetime.GetCurrentMiliseconds(),
// 			UpdatedTime: datetime.GetCurrentMiliseconds(),
// 		}
//	
// 		resultInsert, err := db.Exec("INSERT INTO account (username, password, email, status, role, created_time, updated_time) VALUES (?, ?, ?, ?, ?, ?, ?)",
// 			newAccount.Username, newAccount.Password, newAccount.Email, newAccount.Status, newAccount.Role, newAccount.CreatedTime, newAccount.UpdatedTime)
// 		if err != nil {
// 			logger.Fatal("Failed to insert new account", zap.Error(err))
// 		}
//	
// 		rowInserted, err := resultInsert.RowsAffected()
// 		if err != nil {
// 			logger.Fatal("Failed to insert new account", zap.Error(err))
// 		}
// 		logger.Info("Insert account successed", zap.Int64("Row affected", rowInserted))
//	
// 		//
// 		// Select accounts from DB
// 		var accounts []AccountDTO
// 		if err := db.Select(&accounts, "SELECT * FROM account"); err != nil {
// 			if err != sql.ErrNoRows {
// 				logger.Fatal("Failed to query accounts", zap.Error(err))
// 			}
// 		}
// 		logger.Info("Query accounts", zap.Any("List account", accounts))
//	
// 		//
// 		// Select a account by username
// 		var account AccountDTO
// 		if err := db.Get(&account, "SELECT * FROM account WHERE account.username = ?", "AJPham"); err != nil {
// 			if err == sql.ErrNoRows {
// 				logger.Fatal("Failed to query account", zap.Error(err))
// 			}
// 		}
// 		logger.Info("Query account", zap.Any("Account", account))
//
// 		//
// 		// Update status of account by username
// 		resultUpdate, err := db.Exec("UPDATE account SET account.status = ?, account.updated_time = ? WHERE account.username = ?", 2, datetime.GetCurrentMiliseconds(), "AJPham")
// 		if err != nil {
// 			logger.Fatal("Failed to update account", zap.Error(err))
// 		}
//	
// 		rowUpdated, err := resultUpdate.RowsAffected()
// 		if err != nil {
// 			logger.Fatal("Failed to update new account", zap.Error(err))
// 		}
// 		logger.Info("Update accounts successed", zap.Int64("Row affected", rowUpdated))
// 	}