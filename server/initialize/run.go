package initialize

func RunServer() {
	ConfigInit()
	InitMySQL()
	InitRedis()
	InitAlipay()
	InitRouter()
}
