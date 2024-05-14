package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 读取yam.yml配置文件
type Options struct {
	ListenPort string
	UploadDir  string

	DbHost                    string
	DbPort                    string
	DbUser                    string
	DbPassword                string
	DbDatabase                string
	DbDriver                  string
	DBMaxIdleConn             int
	DBConnectTimeoutInSeconds int
	DBMaxOpenConn             int
}

// 读取iam.yml文件，生成options需要的结果
func NewOption(path string) (*Options, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
		return nil, err
	}

	return &Options{
		ListenPort:                viper.GetString("listen_port"),
		UploadDir:                 viper.GetString("upload_dir"),
		DbPort:                    viper.GetString("db_port"),
		DbUser:                    viper.GetString("db_user"),
		DbHost:                    viper.GetString("db_host"),
		DbPassword:                viper.GetString("db_password"),
		DbDatabase:                viper.GetString("db_database"),
		DbDriver:                  viper.GetString("db_driver"),
		DBMaxIdleConn:             viper.GetInt("db_max_idle_conn"),
		DBConnectTimeoutInSeconds: viper.GetInt("db_connect_timeout_in_seconds"),
		DBMaxOpenConn:             viper.GetInt("db_max_open_conn"),
	}, nil
}

var cfg = &Config{
	Options: &Options{},
}

// Config 自定义配置
type Config struct {
	Options *Options
}

func InitConfig(path string) {
	options, err := NewOption(path)
	if err != nil {
		log.Errorf("解析配置yaml文件失败，错误:[%w]", err)
	}
	cfg.Options = options
}

func GetConfig() *Config {
	return cfg
}
