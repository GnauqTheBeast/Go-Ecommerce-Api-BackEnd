package global

import (
	"github.com/GnauqTheBeast/go-ecommerce-backend-api/pkg/logger"
	"github.com/GnauqTheBeast/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb    *gorm.DB
)
