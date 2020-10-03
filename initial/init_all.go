package initial

import (
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"xxoo/config"
)

func InitializeAll() {
	initMysql()
	initRedis()
}

func initMysql() {
	var mid = config.CONFIG
	var database = mid.Db

	//
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "t_" + defaultTableName
	//}
	//db, err := gorm.Open("mysql", database.Username+":"+database.Password+"@("+database.Path+")/"+database.Dbname+"?"+database.Config)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(database.Username+":"+database.Password+"@tcp("+database.Path+")/"+database.Dbname+"?"+database.Config), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
		Logger:         newLogger,
	},

	)

	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	config.DB = db

	//config.DB.DB().SetMaxIdleConns(database.MaxIdleConns)
	//config.DB.DB().SetMaxOpenConns(database.MaxOpenConns)
	//config.DB.LogMode(database.LogMode)
	//config.DB.SingularTable(true)

}

func initRedis() {
	var mid = config.CONFIG
	var redisCfg = mid.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		println("connect redis err.....")
		return
	}
	config.REDIS = client

}
