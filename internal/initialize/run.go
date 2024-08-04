package initialize

func Run() {
	LoadConfig()
	InitPostgres()

	r := InitRouter()

	r.Run()
}
