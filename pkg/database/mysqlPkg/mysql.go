package mysqlpkg

import (
	"go-clean-arch/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewMysqlDBConnection(cfg *configs.Mysql) (*gorm.DB, error) {
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		defer db.Close()
	}
	return db, err
}
