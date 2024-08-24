package initialize

import (
	"fmt"

	"github.com/GnauqTheBeast/go-ecommerce-backend-api/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() error {
	p := global.Config.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", p.Host, p.Username, p.Password, p.Dbname, p.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	global.Pdb = db

	return nil
}

func SetPool() {

}
