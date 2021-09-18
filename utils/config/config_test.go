package config

import (
	"testing"

	"github.com/phamtai97/go-utils/utils/logger"
	"github.com/stretchr/testify/assert"
)

type ServiceConfigYaml struct {
	Bootstrap  BootstrapConfigYaml
	Datasource DataSourceConfigYaml
}

type BootstrapConfigYaml struct {
	Env            string
	Token          string
	Password       []string
	WorkerPoolSize int  `yaml:"workerPoolSize"`
	EnabledJob     bool `yaml:"enabledJob"`
}

type DataSourceConfigYaml struct {
	AccountDS DatabaseConfigYaml `yaml:"accountDS"`
	SystemDS  DatabaseConfigYaml `yaml:"systemDS"`
}

type DatabaseConfigYaml struct {
	Host      string
	Port      int
	Username  string
	Password  string
	TableName []string `yaml:"tableName"`
}

type ServiceConfigJson struct {
	Bootstrap  BootstrapConfigJson
	Datasource DataSourceConfigJson
}

type BootstrapConfigJson struct {
	Env            string
	Token          string
	Password       []string
	WorkerPoolSize int  `json:"workerPoolSize"`
	EnabledJob     bool `json:"enabledJob"`
}

type DataSourceConfigJson struct {
	AccountDS DatabaseConfigJson `json:"accountDS"`
	SystemDS  DatabaseConfigJson `json:"systemDS"`
}

type DatabaseConfigJson struct {
	Host      string
	Port      int
	Username  string
	Password  string
	TableName []string `json:"tableName"`
}

func TestLoadYaml_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	passwordArr := []string{"abc", "123"}
	tableNameAccArr := []string{"Test1", "Test2", "Test3"}
	tableNameSysArr := []string{"Test4", "Test5", "Test6"}

	// WHEN
	serviceConfig := ServiceConfigYaml{}
	err := LoadYaml(&serviceConfig, "config.yaml")

	// THEN
	assert.Nil(err)
	assert.NotNil(serviceConfig)
	bootstrap := serviceConfig.Bootstrap
	datasource := serviceConfig.Datasource

	assert.NotNil(bootstrap)
	assert.NotNil(datasource)

	assert.Equal("DEV", bootstrap.Env)
	assert.Equal("xyz1234567890", bootstrap.Token)
	assert.Equal(passwordArr, bootstrap.Password)
	assert.Equal(20, bootstrap.WorkerPoolSize)
	assert.Equal(false, bootstrap.EnabledJob)

	accountDS := datasource.AccountDS
	systemDS := datasource.SystemDS

	assert.NotNil(accountDS)
	assert.NotNil(systemDS)

	assert.Equal("9.9.9.9", accountDS.Host)
	assert.Equal(9090, accountDS.Port)
	assert.Equal("ajpham97", accountDS.Username)
	assert.Equal("abc@123", accountDS.Password)
	assert.Equal(tableNameAccArr, accountDS.TableName)

	assert.Equal("8.8.8.8", systemDS.Host)
	assert.Equal(8080, systemDS.Port)
	assert.Equal("ajpham97", systemDS.Username)
	assert.Equal("123@abc", systemDS.Password)
	assert.Equal(tableNameSysArr, systemDS.TableName)
}

func TestLoadYamlByFlag_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	passwordArr := []string{"abc", "123"}
	tableNameAccArr := []string{"Test1", "Test2", "Test3"}
	tableNameSysArr := []string{"Test4", "Test5", "Test6"}

	// WHEN
	serviceConfig := ServiceConfigYaml{}
	err := LoadYamlByFlag(&serviceConfig, "cfgPathYaml")

	// THEN
	assert.Nil(err)
	assert.NotNil(serviceConfig)
	bootstrap := serviceConfig.Bootstrap
	datasource := serviceConfig.Datasource

	assert.NotNil(bootstrap)
	assert.NotNil(datasource)

	assert.Equal("DEV", bootstrap.Env)
	assert.Equal("xyz1234567890", bootstrap.Token)
	assert.Equal(passwordArr, bootstrap.Password)
	assert.Equal(20, bootstrap.WorkerPoolSize)
	assert.Equal(false, bootstrap.EnabledJob)

	accountDS := datasource.AccountDS
	systemDS := datasource.SystemDS

	assert.NotNil(accountDS)
	assert.NotNil(systemDS)

	assert.Equal("9.9.9.9", accountDS.Host)
	assert.Equal(9090, accountDS.Port)
	assert.Equal("ajpham97", accountDS.Username)
	assert.Equal("abc@123", accountDS.Password)
	assert.Equal(tableNameAccArr, accountDS.TableName)

	assert.Equal("8.8.8.8", systemDS.Host)
	assert.Equal(8080, systemDS.Port)
	assert.Equal("ajpham97", systemDS.Username)
	assert.Equal("123@abc", systemDS.Password)
	assert.Equal(tableNameSysArr, systemDS.TableName)
}

func TestLoadYaml_InvalidFile_FailedToLoadConfig(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	serviceConfig := ServiceConfigYaml{}
	err := LoadYaml(&serviceConfig, "./config.txt")

	// THEN
	assert.Equal("open ./config.txt: no such file or directory", err.Error())
}

func TestLoadJson_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	passwordArr := []string{"abc", "123"}
	tableNameAccArr := []string{"Test1", "Test2", "Test3"}
	tableNameSysArr := []string{"Test4", "Test5", "Test6"}

	// WHEN
	serviceConfig := ServiceConfigJson{}
	err := LoadJson(&serviceConfig, "config.json")

	// THEN
	assert.Nil(err)
	assert.NotNil(serviceConfig)
	bootstrap := serviceConfig.Bootstrap
	datasource := serviceConfig.Datasource

	assert.NotNil(bootstrap)
	assert.NotNil(datasource)

	assert.Equal("DEV", bootstrap.Env)
	assert.Equal("xyz1234567890", bootstrap.Token)
	assert.Equal(passwordArr, bootstrap.Password)
	assert.Equal(20, bootstrap.WorkerPoolSize)
	assert.Equal(false, bootstrap.EnabledJob)

	accountDS := datasource.AccountDS
	systemDS := datasource.SystemDS

	assert.NotNil(accountDS)
	assert.NotNil(systemDS)

	assert.Equal("9.9.9.9", accountDS.Host)
	assert.Equal(9090, accountDS.Port)
	assert.Equal("ajpham97", accountDS.Username)
	assert.Equal("abc@123", accountDS.Password)
	assert.Equal(tableNameAccArr, accountDS.TableName)

	assert.Equal("8.8.8.8", systemDS.Host)
	assert.Equal(8080, systemDS.Port)
	assert.Equal("ajpham97", systemDS.Username)
	assert.Equal("123@abc", systemDS.Password)
	assert.Equal(tableNameSysArr, systemDS.TableName)
}

func TestLoadJsonByFlag_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	passwordArr := []string{"abc", "123"}
	tableNameAccArr := []string{"Test1", "Test2", "Test3"}
	tableNameSysArr := []string{"Test4", "Test5", "Test6"}

	// WHEN
	serviceConfig := ServiceConfigJson{}
	err := LoadJsonlByFlag(&serviceConfig, "cfgPathJson")

	// THEN
	assert.Nil(err)
	assert.NotNil(serviceConfig)
	bootstrap := serviceConfig.Bootstrap
	datasource := serviceConfig.Datasource

	assert.NotNil(bootstrap)
	assert.NotNil(datasource)

	assert.Equal("DEV", bootstrap.Env)
	assert.Equal("xyz1234567890", bootstrap.Token)
	assert.Equal(passwordArr, bootstrap.Password)
	assert.Equal(20, bootstrap.WorkerPoolSize)
	assert.Equal(false, bootstrap.EnabledJob)

	accountDS := datasource.AccountDS
	systemDS := datasource.SystemDS

	assert.NotNil(accountDS)
	assert.NotNil(systemDS)

	assert.Equal("9.9.9.9", accountDS.Host)
	assert.Equal(9090, accountDS.Port)
	assert.Equal("ajpham97", accountDS.Username)
	assert.Equal("abc@123", accountDS.Password)
	assert.Equal(tableNameAccArr, accountDS.TableName)

	assert.Equal("8.8.8.8", systemDS.Host)
	assert.Equal(8080, systemDS.Port)
	assert.Equal("ajpham97", systemDS.Username)
	assert.Equal("123@abc", systemDS.Password)
	assert.Equal(tableNameSysArr, systemDS.TableName)
}

func TestLoadJson_InvalidFile_FailedToLoadConfig(t *testing.T) {
	// GIVEN
	assert := assert.New(t)

	// WHEN
	serviceConfig := ServiceConfigJson{}
	err := LoadJson(&serviceConfig, "./config.txt")

	// THEN
	assert.Equal("open ./config.txt: no such file or directory", err.Error())
}

func TestPrint_SimpleInput_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	logger.InitProduction("")

	// WHEN
	serviceConfig := ServiceConfigYaml{}
	err := LoadYaml(&serviceConfig, "config.yaml")
	errPrint := Print(serviceConfig)

	// THEN
	assert.Nil(err)
	assert.Nil(errPrint)
}
func TestPrint_OmittedKeys_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	logger.InitProduction("")

	// WHEN
	serviceConfig := ServiceConfigYaml{}
	err := LoadYaml(&serviceConfig, "config.yaml")
	errPrint := Print(serviceConfig, "Token", "Password")

	// THEN
	assert.Nil(err)
	assert.Nil(errPrint)
}

func TestPrint_ConfigIsNull_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	logger.InitProduction("")

	// WHEN
	errPrint := Print(nil, "Token", "Password")

	// THEN
	assert.Nil(errPrint)
}

func TestPrint_InvalidConfig_Success(t *testing.T) {
	// GIVEN
	assert := assert.New(t)
	logger.InitProduction("")

	// WHEN
	errPrint := Print("", "Token", "Password")

	// THEN
	assert.Equal("json: cannot unmarshal string into Go value of type map[string]interface {}", errPrint.Error())
}

func BenchmarkLoadYaml(b *testing.B) {
	serviceConfig := ServiceConfigYaml{}
	LoadYaml(&serviceConfig, "config.yaml")
}

func BenchmarkLoadJson(b *testing.B) {
	serviceConfig := ServiceConfigJson{}
	LoadJson(&serviceConfig, "config.yaml")
}
