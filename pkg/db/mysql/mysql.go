package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	rawDB     *gorm.DB
	debugMode bool
)

type DB struct {
	db                *gorm.DB      // 数据库实例
	transactionEnable bool          // 事务开启状态
	logger            *logrus.Entry // 日志
	Ctx               *gin.Context  // gin context
	*User                           // 当前用户状态
}

type User struct {
	TokenUUid   string
	IsLogin     bool
	UserID      string
	UserName    string
	DisplayName string
}

func NewDB() *DB {
	if debugMode {
		return &DB{
			db: rawDB.Debug(),
		}
	}
	return &DB{
		db: rawDB,
	}
}

func NewTX() *DB {
	if debugMode {
		return &DB{
			db: rawDB.Debug().Begin(),
		}
	}
	return &DB{
		db: rawDB.Begin(),
	}
}

func (b *DB) GetDB() *gorm.DB {
	if b.db == nil {
		return rawDB
	}
	return b.db
}

func (b *DB) SetDB(db *gorm.DB) {
	b.db = db
}

func (b *DB) Begin() *gorm.DB {
	b.db = rawDB.Begin()
	return b.db
}

func (b *DB) Commit() {
	if b.db != nil {
		b.db.Commit()
	}
}

type MySQLConfig struct {
	DbHost                    string
	DbPort                    string
	DbUser                    string
	DbPassword                string
	DbDatabase                string
	DBDriver                  string
	DBMaxIdleConn             int
	DBConnectTimeoutInSeconds int
	DBMaxOpenConn             int
}

func InitMysql(cfg *MySQLConfig) (err error) {
	var dialector gorm.Dialector
	if cfg.DBDriver == "mysql" {
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbDatabase))
	}

	rawDB, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	rawDB.Exec("set time_zone=\"+08:00\";")
	if err != nil {
		// 数据库未创建
		if strings.Contains(err.Error(), "Unknown database") {
			logrus.Fatalf("数据库[%s] 未创建", cfg.DbDatabase)
		}
		// 数据库连接失败
		logrus.Fatalf("Connect database failed, err: %v", err)
		return err
	}
	sqlDB, _ := rawDB.DB()
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnectTimeoutInSeconds))
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConn)
	return nil
}
