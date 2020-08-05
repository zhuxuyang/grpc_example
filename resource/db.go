package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log"
	"time"
)

var Db *gorm.DB

func InitDB() {
	dbConf := viper.GetStringMapString("database")
	initDb(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])
}
func initDb(user, password, host, port, dbName string) {
	log.Println("connecting MySQL ... ", host)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	mdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &Mylog{ServiceLog: *Logger},
	})
	if mdb == nil {
		log.Println("failed to connect database")
		return
	}
	log.Println("connected")
	Db = mdb
	return
}

type Mylog struct {
	ServiceLog Loggers
}

// LogMode log mode
func (l *Mylog) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

//
//// Info print info
func (l Mylog) Info(ctx context.Context, msg string, data ...interface{}) {
	l.ServiceLog.Info(msg, data)
	//if l.LogLevel >= Info {
	//	l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	//}
}

// Warn print warn messages
func (l Mylog) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.ServiceLog.Warn(msg, data)
	//if l.LogLevel >= Warn {
	//	l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	//}
}

// Error print error messages
func (l Mylog) Error(ctx context.Context, msg string, data ...interface{}) {
	l.ServiceLog.Error(msg, data)
	//if l.LogLevel >= Error {
	//	l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	//}
}

// Trace print sql message
func (l Mylog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Now().Sub(begin)
	s, _ := json.Marshal(&ctx)
	l.Info(ctx, string(s))
	if err != nil {
		sql, rows := fc()
		l.ServiceLog.Error(ctx, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	} else {
		sql, rows := fc()
		l.ServiceLog.Info(utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
