package bootstrap

import (
	"encoding/json"
	"fmt"
	"ginapi/app/global"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// 数据库配置类型
type DBConfig struct {
	Name     string
	WriteDSN string
	ReadDSNs []string
}

func InitializeDB() map[string]*gorm.DB {

	// fmt.Println(global.App.Config.Db)
	// 暂时写死，稍后改为配置文件读取
	// dbHostWrite := "127.0.0.1"
	dbHostWrite := global.App.Config.Db.DbHostWrite
	// fmt.Println(dbHostWrite)
	// dbHostRead := "127.0.0.1,127.0.0.1"
	dbHostRead := global.App.Config.Db.DbHostRead

	// dbPortWrite := "3306"
	dbPortWriteInt := global.App.Config.Db.DbPortWrite
	dbPortWrite := fmt.Sprintf("%d", dbPortWriteInt)
	// fmt.Println("dbPortWrite:", dbPortWrite)
	// dbPortRead := "3306"
	dbPortReadInt := global.App.Config.Db.DbPortRead
	// dbPortRead把int转字符串类型
	dbPortRead := fmt.Sprintf("%d", dbPortReadInt)

	// fmt.Println("dbPortRead:", dbPortRead)
	// dbUser := "xxxxx"
	dbUser := global.App.Config.Db.UserName

	// dbPassword := "xxxxxxx"
	dbPassword := global.App.Config.Db.Password

	queryParams := "charset=utf8mb4&parseTime=True&loc=Local"

	// 配置需要连接的数据库,暂时写死，跑通后改为配置文件读取
	// dbNameList := []string{"ppc_library", "growth"}
	dbNameList := strings.Split(global.App.Config.Db.DbNameList, ",")

	// 创建一个 map 来保存所有的数据库连接
	var dbs = make(map[string]*gorm.DB)

	// dbHostRead 按','分割成切片
	dbHostReads := strings.Split(dbHostRead, ",")
	// fmt.Println("dbHostReads:", dbHostReads)
	// 声明读库切片
	var readDSNs []string

	for _, dbName := range dbNameList {
		// fmt.Println("dbName:", dbName)
		writeDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPassword, dbHostWrite, dbPortWrite, dbName, queryParams)

		for _, readHost := range dbHostReads {
			readDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPassword, readHost, dbPortRead, dbName, queryParams)
			readDSNs = append(readDSNs, readDSN)
		}

		DBConfig := DBConfig{
			Name:     dbName,
			WriteDSN: writeDSN,
			ReadDSNs: readDSNs,
		}

		readDSNs = nil

		dbs[DBConfig.Name] = createDBResolver(DBConfig)
	}
	return dbs
}

// 创建数据库连接配置
func createDBResolver(config DBConfig) *gorm.DB {

	// 打印配置信息
	a, _ := json.Marshal(config)
	str := string(a)
	fmt.Println(str)

	// 主库连接
	writeConn := mysql.Open(config.WriteDSN)

	// 从库连接
	readConns := make([]gorm.Dialector, len(config.ReadDSNs))
	for i, readDSN := range config.ReadDSNs {
		readConns[i] = mysql.Open(readDSN)
	}

	// 打开数据库连接
	db, err := gorm.Open(writeConn, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
	})
	if err != nil {
		log.Fatalf("failed to connect database %s: %v", config.Name, err)
	}

	// 注册读写分离规则
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{writeConn},
		Replicas: readConns,
		Policy:   dbresolver.RandomPolicy{}, // 使用随机策略选择副本
	}))

	return db
}
