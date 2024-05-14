package app

import (
	"github.com/Touch/cmd/api/routers"
	"github.com/Touch/config"
	"github.com/Touch/pkg/db/mysql"
	"github.com/gin-gonic/gin"
)

func Run() {
	mysqlConfig := mysql.MySQLConfig{
		DbPort:                    config.GetConfig().Options.DbPort,
		DbUser:                    config.GetConfig().Options.DbUser,
		DbHost:                    config.GetConfig().Options.DbHost,
		DbPassword:                config.GetConfig().Options.DbPassword,
		DbDatabase:                config.GetConfig().Options.DbDatabase,
		DBDriver:                  config.GetConfig().Options.DbDriver,
		DBMaxIdleConn:             config.GetConfig().Options.DBMaxIdleConn,
		DBConnectTimeoutInSeconds: config.GetConfig().Options.DBConnectTimeoutInSeconds,
		DBMaxOpenConn:             config.GetConfig().Options.DBMaxOpenConn,
	}
	mysql.InitMysql(&mysqlConfig)

	r := gin.Default()
	routers.RegisterRoutes(r)
	// live probe
	r.Run(":8006")
}
