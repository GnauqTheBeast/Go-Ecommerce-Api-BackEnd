package initialize

import (
	"github.com/GnauqTheBeast/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	err := LoadConfig()
	if err != nil {
		global.Logger.Error("Loading config file", zap.String("load config", "no"))
		return
	}

	InitLogger()
	global.Logger.Info("Loading logger", zap.String("load logger", "yes"))

	err = InitPostgres()
	if err != nil {
		global.Logger.Error("Loading db", zap.String("db", "no"))
		return
	}
	global.Logger.Info("Loading DB", zap.String("load DB", "yes"))

	r := InitRouter()

	r.Run()
}

func checkError(err error, errString string) {

}
