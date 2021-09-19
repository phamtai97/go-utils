// Package config provides a way to load configuration of application from yaml and json file.
//
// Example Usage
//
// The following is a complete example using logger package
// 	import (
// 		"github.com/phamtai97/go-utils/utils/config"
// 		"github.com/phamtai97/go-utils/utils/logger"
// 		"go.uber.org/zap"
// 	)
//
// 	// ServiceConfig config to test
// 	type ServiceConfig struct {
// 		Bootstrap  BootstrapConfig
// 		Datasource DataSourceConfig
// 	}
//
// 	// BootstrapConfig config to test
// 	type BootstrapConfig struct {
// 		Env            string
// 		Token          string
// 		Password       []string
// 		WorkerPoolSize int  `yaml:"workerPoolSize"`
// 		EnabledJob     bool `yaml:"enabledJob"`
// 	}
//
// 	// DataSourceConfig config to test
// 	type DataSourceConfig struct {
// 		AccountDS DatabaseConfig `yaml:"accountDS"`
// 		SystemDS  DatabaseConfig `yaml:"systemDS"`
// 	}
//
// 	// DatabaseConfig config to test
// 	type DatabaseConfig struct {
// 		Host      string
// 		Port      int
// 		Username  string
// 		Password  string
// 		TableName []string `yaml:"tableName"`
// 	}
//
// 	func main() {
// 		logger.InitProduction("")
// 		serviceConfig := ServiceConfig{}
//
// 		// We can provide path of config by flag to load config
// 		// config.LoadByFlag(&serviceConfig, "cfgPath")
// 		if err := config.Load(&serviceConfig, "dev.yaml"); err != nil {
// 			logger.Fatal("Failed to load config", zap.Error(err))
// 		}
//
// 		// Load config from json file
// 		// if err := config.Load(&serviceConfig, "dev.json"); err != nil {
// 		// 	logger.Fatal("Failed to load config", zap.Error(err))
// 		// }
//
// 		if err := config.Print(serviceConfig, "Token", "Password"); err != nil {
// 			logger.Fatal("Failed to print config", zap.Error(err))
// 		}
// 	}
package config
