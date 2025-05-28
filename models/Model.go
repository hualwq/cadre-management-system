package models

import (
	"cadre-management/pkg/setting"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)

	// 2. 初始化方式变化（需显式指定驱动）
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 3. 表名前缀处理（通过 NamingStrategy 配置）
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix, // 直接配置前缀
			SingularTable: true,                                // 等效于旧版的 db.SingularTable(true)
		},
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// 4. 回调注册方式变化（使用 CreateHook 等接口） 在创建、更新、删除数据之前自动创建时间戳 create_at, update_at, delete_at
	// err = db.Callback().Create().Before("gorm:create").Register("update_time_stamp", updateTimeStampForCreateCallback)
	// err = db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimeStampForUpdateCallback)
	// err = db.Callback().Delete().Before("gorm:delete").Register("delete_hook", deleteCallback)
	if err != nil {
		log.Fatalf("failed to register callbacks: %v", err)
	}

	// 5. 连接池配置（通过 sql.DB 对象）
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func GetDbInstance() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDB 关闭数据库连接
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to get underlying sql.DB: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		var extraOption string
		if str, ok := db.Statement.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if field := db.Statement.Schema.LookUpField("DeletedOn"); field != nil && !db.Statement.Unscoped {

			db.Exec(fmt.Sprintf(
				"UPDATE %s SET %s = ? %s %s",
				db.Statement.Table,
				field.DBName,
				db.Statement.SQL.String(), // WHERE 条件
				extraOption,
			), time.Now().Unix())
		} else {
			// 硬删除逻辑
			db.Exec(fmt.Sprintf(
				"DELETE FROM %s %s %s",
				db.Statement.Table,
				db.Statement.SQL.String(),
				extraOption,
			))
		}
	}
}

// addExtraSpaceIfExist 添加分隔符
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
